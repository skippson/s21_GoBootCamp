package main

import (
	"flag"
	"fmt"
	"log"
	reader "secondDay/reader"
	"strings"
)

func main() {
	filename := flag.String("f", "", "input file name")
	flag.Parse()
	if *filename == "" {
		log.Fatal("no input file specified")
	}
	ras := strings.HasSuffix(*filename, ".json")
	db := reader.DB{}
	if !ras {
		data, _, err := db.DBReader(*filename, "xml")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s", data)
	} else if ras {
		data, _, err := db.DBReader(*filename, "json")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s", data)

	} else {
		fmt.Println("invalid filename")
	}
}
