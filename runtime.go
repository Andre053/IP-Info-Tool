package main

import (
	"errors"
	"fmt"
)

func printIntro() {
	fmt.Println("Welcome to my IP Lookup tool")
}

func printOptions() {
	fmt.Printf("\n1 - Geo Lookup\t\t2 - ASN Lookup\t\t3 - Distance Lookup\t\t4 - Exit\n\n")
}

func run() {
	var choice string
	cache := createCache()
	printIntro()
	for {
		printOptions()
		fmt.Printf("Select an option: ")
		_, err := fmt.Scanln(&choice)

		success, err := handleChoice(choice, cache)
		if err != nil {
			fmt.Println(err)
			break
		}
		if !success {
			fmt.Println("Failed, try again")
		}
	}
	fmt.Println("Exiting program...")
}

func handleChoice(choice string, c *cache) (bool, error) {
	switch choice {
	case "1": // geo
		return true, handleIpGeo(c)
	case "2": // asn
		return true, handleIpAsn(c)
	case "3": // distance
		fmt.Println("Handling distance lookup")
		return true, nil
	case "4": // exit
		return false, errors.New("Quit program")
	default:
		return false, errors.New("Unsupported choice")
	}
}

func getIpOrHostnameFromUser() (string, error) {
	var input string
	fmt.Println("Grabbing IP from user")
	return input, nil
}
