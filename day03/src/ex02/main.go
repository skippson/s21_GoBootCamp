package main

import (
	"day03/ex02/api"
	"log"

	"github.com/elastic/go-elasticsearch/v8"
)

func main() {
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatal(err)
	}
	a := api.API{}
	a.ApiShow(es, 13649)

}
