package main

import (
	"bytes"
	"encoding/json"
	"io"
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
	token := os.Getenv("IP_INFO_TOKEN")
	url := "https://ipinfo.io/" + ip.String() + "?" + token

	res, err := handleHTTPRequest(url)
	if err != nil {
		return *res, err
	}
	return *res, nil
}

func handleHTTPRequest(url string) (*Response, error) {

	var data *Response
	dataBytes := make([]byte, 1024)

	res, err := http.Get(url)
	if err != nil {
		return &Response{}, err
	}

	read, err := res.Body.Read(dataBytes)
	if err != nil {
		if err != io.EOF {
			return &Response{}, err
		}
	}
	if read == 0 {
		return &Response{}, &NoHttpData{Line: 43, Col: 4}
	}
	dataBytes = bytes.Trim(dataBytes, "\x00")

	err = json.Unmarshal(dataBytes, &data)
	if err != nil {
		return &Response{}, err
	}
	if data.Ip == "" || data.Hostname == "" || data.City == "" || data.Region == "" || data.Country == "" || data.Loc == "" || data.Org == "" || data.Postal == "" || data.Timezone == "" {
		return &Response{}, &MissingJsonValuesError{Line: 51, Col: 4}
	}
	return data, nil
}
