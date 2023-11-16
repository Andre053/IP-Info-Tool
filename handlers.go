package main

import (
	"bytes"
	"encoding/json"
	"net"
	"net/http"
	"os"
)

func getIpInfo(ip net.IP) (Response, error) {

	res, err := handleLookup(ip)
	if err != nil {
		return res, err
	}

	return res, nil
}

func handleLookup(ip net.IP) (Response, error) {
	res, err := handleHTTPRequest(ip)
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
