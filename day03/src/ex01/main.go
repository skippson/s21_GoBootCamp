package main

import (
	"day03/ex01/http"
	"log"

	"github.com/elastic/go-elasticsearch/v8"
)

func main() {
	es, err := elasticsearch.NewDefaultClient()
	
	if err != nil {
		log.Fatal(err)
	}

	hp := http.HTTP{}
	hp.HttpShow(es)
}
