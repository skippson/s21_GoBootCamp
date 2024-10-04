package main

import (
	"fmt"
	."day09/ex01/spider"
)

func main() {
	done := make(chan struct{})
	WatchSigterm(done)

	url := make(chan string)
	go SendURSs(url, done)

	data := CrawlWeb(url, done)
	for val := range data {
		fmt.Println(val)
	}
}
