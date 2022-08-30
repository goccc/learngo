package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	startTime := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch)

	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(startTime).Seconds())

}

func fetch(url string, ch chan<- string) {
	startTime := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(io.Discard, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		resp.Body.Close()
		return
	}
	resp.Body.Close()
	secs := time.Since(startTime).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}
