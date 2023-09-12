package main

import (
	"golang.org/x/net/publicsuffix"
	"log"
)

func main() {

	sld, err := publicsuffix.EffectiveTLDPlusOne("www.httpbin.org")
	if err == nil {
		log.Println(sld)
	}

}
