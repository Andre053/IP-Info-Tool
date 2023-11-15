# DESIGN OVERVIEW
- CLI for querying IP addresses
- Allows for IP or domain name
- Gathers info about the IP, such as geo location

# CLI
## General Mode
1. Choose an action
2. Choose an IP address or hostname
3. See info
4. Repeat

## All Info Mode
1. Choose an IP/domain name
2. Gathers and displays all info related to it
3. Repeat

# API
- Utilizes ipinfo.io

# Improving Performance
- Added a software cache
    - Will cache requests to IPs
    - Additional data structure adds much overhead
    - Much faster if querying the same IP address again
    - NOTE: Mainly for proof of concept

# Testing
- Most tests will use lookback server 