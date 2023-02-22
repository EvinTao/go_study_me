package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	defer func() {
		if resp != nil {
			resp.Body.Close() //不要泄露资源
		}
	}()
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(io.Discard, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	sesc := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", sesc, nbytes, url)

}

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) //启动一个goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) //从通道ch接收
	}
	fmt.Printf("%2.fs elapsed \n", time.Since(start).Seconds())
}
