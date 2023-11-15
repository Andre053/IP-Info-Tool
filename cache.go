package main

import "fmt"

type cache struct {
	lookups map[string]Response
}

func createCache() *cache {
	return &cache{
		lookups: map[string]Response{},
	}
}

func (c *cache) addToCache(r Response) error {
	c.lookups[r.Ip] = r

	return nil
}

func (c cache) checkCache(ip string) (Response, bool) {

	res, found := c.lookups[ip]
	if !found {
		return res, false
	}
	fmt.Println("[DEBUG] Found in cache!")
	return res, true
}
