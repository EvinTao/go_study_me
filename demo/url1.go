package main

import (
	"fmt"
	"net/url"
)

func main() {
	urlString := "http://kdlime:GizaPk3V2F@113.251.72.174:55001"
	parsedURL, err := url.Parse(urlString)

	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return
	}

	fmt.Println(parsedURL)
}
