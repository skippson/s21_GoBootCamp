package spider

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func WatchSigterm(done chan struct{}) {
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		close(done)
	}()
}

func SendURSs(ch chan<- string, done chan struct{}) {
	for i := 11; i < 50; i++ {
		time.Sleep(1 * time.Second)
		url := "https://api.disneyapi.dev/character/" + fmt.Sprintf("%d", i)
		select {
		case ch <- url:
		case <-done:
			close(ch)
			os.Exit(0)
		}
	}

	close(ch)
}

func CrawlWeb(ch <-chan string, done chan struct{}) chan string {
	data := make(chan string)
	maxim := make(chan struct{}, 8)

	go func() {
		wg := sync.WaitGroup{}
		for url := range ch {
			maxim <- struct{}{}
			wg.Add(1)
			go func() {
				getData(url, data, &wg, done)
				<-maxim
			}()
		}

		wg.Wait()
		close(data)
	}()

	return data
}

func getData(url string, ch chan string, wg *sync.WaitGroup, done chan struct{}) {
	defer wg.Done()

	r, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	select {
	case ch <- string(body):
	case <-done:
		close(ch)
		os.Exit(0)
	}
}
