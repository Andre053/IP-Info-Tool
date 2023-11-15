package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
)

func getIpsFromUser() ([]net.IP, error) {
	var input string
	var ips []net.IP
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
			ips = append(ips, ip)
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
		ips = append(ips, ipsFound[0])
	}
	return ips, nil
}
func handleIpGeo(c *cache) error {
	ips, err := getIpsFromUser()
	if err != nil {
		return err
	}
	for _, ip := range ips {
		info, err := getIpGeo(ip, c)
		if err != nil {
			return err
		}
		fmt.Println(info)
	}
	return nil
}

func handleIpAsn(c *cache) error {
	ips, err := getIpsFromUser()
	if err != nil {
		return err
	}
	for _, ip := range ips {
		info, err := getIpAsnOrg(ip, c)
		if err != nil {
			return err
		}
		fmt.Println(info)
	}
	return nil
}

func handleLookup(ip net.IP, c *cache) (Response, error) {
	cacheRes, found := c.checkCache(ip.String())
	if found {
		return cacheRes, nil
	}
	res, err := handleHTTPRequest(ip)
	if err != nil {
		return *res, err
	}
	err = c.addToCache(*res)
	if err != nil {
		return *res, err
	}
	return *res, nil
}

func handleHTTPRequest(ip net.IP) (*Response, error) {
	token := os.Getenv("ipinfo_TOKEN")
	url := "https://ipinfo.io/" + ip.String() + "?" + token

	var data *Response
	dataBytes := make([]byte, 1024)

	res, err := http.Get(url)
	if err != nil {
		return data, err
	}

	_, err = res.Body.Read(dataBytes)
	if err != nil {
		return data, err
	}
	dataBytes = bytes.Trim(dataBytes, "\x00")

	err = json.Unmarshal(dataBytes, &data)
	if err != nil {
		return data, err
	}
	return data, nil

}
