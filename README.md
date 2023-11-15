# IP Info Lookup Tool v1

A Go CLI for retrieving information about IP addresses.

Does not utilize a library for CLI.

## Usage

Install the code:
`go install github.com/andre053/ip-info-tool`

Get your own API key from https://ipinfo.io/signup

Save your key as an environment variable named "IP_INFO_TOKEN"

Utilize make to interact with the code:

`make build` will build the module into an ipinfo executable.  
`make run` will build and run the ipinfo executable.  
`make clean` will clear the console, run go clean, and remove the executable.  

Add the program to your path:   `export PATH=$PATH:<path to ipinfo>`  

## Features

- Geolocation lookup: Get the city, country code, and coordinates
- ASN lookup: Get the hostname and organization


## In Progress

- Testing suite, utilizing loopback interfaces
- Distance lookup: Calculating distance between IP addresses
- Simple lookups: Domain to IP, IP to domain