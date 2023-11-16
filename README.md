# IP Network Tool

A Go CLI for completing tasks related to IP addresses

## Usage

Install the code:
`go install github.com/andre053/ip-info-tool`

Get your own API key from https://ipinfo.io/signup

Save your key as an environment variable named "IP_INFO_TOKEN"

Utilize make to interact with the code:

`make build` will build the module into an ipinfo executable.  
`make run` will build and run the ipinfo executable, in the new mode, this will call Usage immediately as there are no flags  
`make clean` will clear the console, run go clean, and remove the executable.  

Add the program to your path:   `export PATH=$PATH:<path to ipinfo>`  

## 3rd Party Libraries
- github.com/urfave/cli/v2                  Utilized to simplify generation of CLI commands and usage information
- github.com/prometheus-community/pro-bing  Used to implement ping    
## Features
- CLI: Accepts commands and command shortcuts, queries user for data (no flags for data)
- IP Info: Gathers information related to an IP address from ipinfo.io
- Distance: Calculates the distance between two IP addresses in kilometers
- Ping: Pings the address, returning the statistics
    - Privileged ping requires privileges but will work on systems which do not support unprivileged ping
- Interprets user input: Users can input hostnames of IP addresses, they will be converted by the program
- Multiple inputs: Users can supply a list of inputs for the program to execute on
    - NOTE: Distance only works on two inputs and ping is configured to timeout after 10 seconds

## In Progress

- Testing suite, utilizing loopback interfaces
- Simple lookups: Domain to IP, IP to domain
- Traceroute implementation
- Add flags for commands
- Improve documentation of commands

## Future Challenges

- Implement ping natively
