package main

import (
	"fmt"
	"math"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Response struct {
	Ip       string `json:"ip"`
	Hostname string `json:"hostname"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"loc"`
	Org      string `json:"org"`
	Postal   string `json:"postal"`
	Timezone string `json:"timezone"`
}

type ipIds struct {
	ip       net.IP
	hostname string
}

func getIpsFromUser() ([]ipIds, error) {
	var input string
	var reqIpIds []ipIds

	fmt.Println("Please enter IP addresses or hostnames to lookup, separated by commas:")

	_, err := fmt.Scanln(&input)
	if err != nil {
		return nil, err
	}
	fmt.Println()
	inputs := strings.Split(input, ",")
	for _, val := range inputs {
		ip := net.ParseIP(val)
		if ip != nil {
			// add to results,
			hostnames, err := net.LookupAddr(val)
			if err != nil {
				return reqIpIds, err
			}
			reqIpIds = append(reqIpIds, ipIds{
				ip:       ip,
				hostname: hostnames[0],
			})

		}
		ipsFound, err := net.LookupIP(val)
		if err != nil {
			return nil, err
		}

		// no matching ips found
		if len(ipsFound) < 1 {
			fmt.Printf("No matching IPs for %s\n", val)
			continue
		}
		// add the first IP of the hostname found
		reqIpIds = append(reqIpIds, ipIds{
			ip:       ipsFound[0],
			hostname: val,
		})
	}
	return reqIpIds, nil
}

func mustStrToFloat(s string) float64 {
	f, err := strconv.ParseFloat(s, 4)
	if err != nil {
		panic(err)
	}
	return f
}

func degToRad(d float64) float64 {
	return d * math.Pi / 180
}
func harversineDist(x1, y1, x2, y2 float64) float64 {
	worldRad := float64(6371000)
	lat1 := degToRad(x1)
	lon1 := degToRad(y1)
	lat2 := degToRad(x2)
	lon2 := degToRad(y2)

	latDiff := lat2 - lat1
	lonDiff := lon2 - lon1

	// formula
	a := math.Abs(math.Sin(latDiff/2.0)*2.0 + math.Cos(lat1)*math.Cos(lat2)*math.Sin(lonDiff/2.0)*2.0)
	c := 2.0 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	d := worldRad * c

	return d / 1000.0
}

func waitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	done := make(chan struct{})

	go func() {
		defer close(done)
		wg.Wait()
	}()

	select {
	case <-done:
		return false
	case <-time.After(timeout):
		return true
	}
}
