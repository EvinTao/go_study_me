package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	url := "https://www.baidu.com"
	concurrency := 10
	results := make(chan string)

	for i := 0; i < concurrency; i++ {
		go func() {
			for {
				resp, err := http.Get(url)
				if err != nil {
					fmt.Println(err)
					continue
				}

				doc, err := goquery.NewDocumentFromReader(resp.Body)
				if err != nil {
					fmt.Println(err)
					continue
				}

				title := doc.Find("title").Text()
				results <- title

				time.Sleep(1 * time.Second)
			}
		}()
	}

	for i := 0; i < concurrency; i++ {
		fmt.Println(<-results)
	}
}
