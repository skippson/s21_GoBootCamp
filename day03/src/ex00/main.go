package main

import (
	"log"
	"day03/types"
	"day03/ex00/base"
	"day03/ex00/parser"
)

func main() {
	log.SetFlags(0)
	db := base.Data{}
	p := parser.Pars{}
	var	places []*types.Place
	places, err := p.ParseCSV(places)
	if err != nil {
		log.Fatal(err)
	}
	db.CreateESC(places)
}
