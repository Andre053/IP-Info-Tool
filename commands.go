package main

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	probing "github.com/prometheus-community/pro-bing"
)

func ipInfo() error {

	var data string
	ips, err := getIpsFromUser()
	if err != nil {
		return err
	}

	for _, v := range ips {
		res, err := getIpInfo(v.ip)
		if err != nil {
			return err
		}
		data += "IP Address:\t\t" + res.Ip + "\nHostname:\t\t" + res.Hostname + "\nOrganization:\t\t" + res.Org + "\nCity:\t\t" + res.City + "\nCountry:\t" + res.Country + "\nCoordinates:\t" + res.Loc + "\n"
	}

	fmt.Println(data)
	return nil
}

func ipDistance() error {
	ips, err := getIpsFromUser()
	if err != nil {
		return err
	}
	if len(ips) != 2 {
		return errors.New("IP distance requires two arguments")
	}
	res1, err := getIpInfo(ips[0].ip)
	if err != nil {
		return err
	}
	res2, err := getIpInfo(ips[1].ip)
	if err != nil {
		return err
	}
	s1 := strings.Split(res1.Loc, ",")
	s2 := strings.Split(res2.Loc, ",")

	x1 := mustStrToFloat(s1[0])
	y1 := mustStrToFloat(s1[1])
	x2 := mustStrToFloat(s2[0])
	y2 := mustStrToFloat(s2[1])

	distance := harversineDist(x1, y1, x2, y2)
	fmt.Printf("Distance between %s (%s) and %s (%s) is %.2fkm\n", res1.Hostname, res1.Ip, res2.Hostname, res2.Ip, distance)

	return nil
}

func ping(privileged bool) error {
	var wg sync.WaitGroup
	PING_COUNT := 5
	ips, err := getIpsFromUser()
	if err != nil {
		return err
	}

	for _, v := range ips {

		wg.Add(1)
		go func(ip, host string) {
			defer wg.Done()

			pinger, err := probing.NewPinger(ip)
			if err != nil {
				panic(err)
			}
			if privileged {
				pinger.SetPrivileged(true)
			}
			pinger.Count = PING_COUNT
			fmt.Println("Pinging", ip)
			err = pinger.Run() // blocking
			if err != nil {
				panic(err)
			}
			stats := pinger.Statistics()
			fmt.Printf("Recieved %d/%d packets from %s (%s), resulting in %.1f%% packet loss, average RTT was %d microseconds\n", stats.PacketsRecv, stats.PacketsSent, host, stats.Addr, stats.PacketLoss, stats.AvgRtt.Microseconds())
		}(v.ip.String(), v.hostname)
	}
	if waitTimeout(&wg, time.Second*10) { // if any pings take longer than 5 seconds, kill it
		return errors.New("Pinging host timed out")
	}

	return nil
}
