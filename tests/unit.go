package main

import (
	"net/http/httptest"
	"testing"
)

type Tests struct {
	name          string
	server        *httptest.Server
	response      *Response
	expectedError error
}

func TestLookupDNS(t *testing.T) {

}

func TestLookupReverseDNS(t *testing.T) {
	// use loopback

}

func TestGetIpGeo(t *testing.T) {

}

func TestGetIpASN(t *testing.T) {

}

func TestGetIpAbuse(t *testing.T) {

}

func TestGetDistanceBetweenIps(t *testing.T) {

}
