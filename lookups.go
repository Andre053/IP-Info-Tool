package main

import (
	"fmt"
	"net"
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

func getIpGeo(ip net.IP, c *cache) (string, error) {
	var geoData string

	res, err := handleLookup(ip, c)
	if err != nil {
		return "", err
	}
	geoData = "City:\t\t" + res.City + "\nCountry:\t" + res.Country + "\nCoordinates:\t" + res.Loc

	return geoData, nil
}

func getIpAsnOrg(ip net.IP, c *cache) (string, error) {
	var asnData string

	res, err := handleLookup(ip, c)
	if err != nil {
		return "", err
	}
	asnData = "Hostname:\t" + res.Hostname + "\nOrganization:\t" + res.Org

	return asnData, nil
}

func getDistanceBetweenIps(ip1 net.IP, ip2 net.IP, c *cache) (string, error) {
	fmt.Printf("Computing data between %v and %v", ip1, ip2)
	return "distance", nil
}
