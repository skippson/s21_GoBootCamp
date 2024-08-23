package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

type flags struct {
	w bool
	l bool
	m bool
}

func main() {
	flags := flags{}
	flags.parse()

	files := getFiles()
	wc(files, flags)

}

func (f *flags) parse() *flags {
	flag.BoolVar(&f.w, "w", false, "menu")
	flag.BoolVar(&f.l, "l", false, "menu")
	flag.BoolVar(&f.m, "m", false, "menu")
	flag.Parse()

	if (f.w && f.l) || (f.w && f.m) || (f.m && f.l) {
		log.Fatal("Use only one flag [-m -l -w]")
	}

	if !f.l && !f.m && !f.w {
		f.w = true
	}

	return f
}

func getFiles() []string {
	arg := os.Args[1:]
	var files []string

	for _, file := range arg {

		if file[0] == '-' {
			continue
		} else {
			files = append(files, file)
		}

	}

	return files
}

func wc(files []string, f flags) {
	type pair struct {
		filename string
		num      int
	}

	result := make([]pair, len(files))

	var wg sync.WaitGroup

	wg.Add(len(files))

	for i, file := range files {
		result[i].filename = file
		go func() {
			defer wg.Done()
			temple, err := os.Open(file)
			if err != nil {
				log.Fatal("Invalid file", temple)
			}
			scanner := bufio.NewScanner(temple)
			for scanner.Scan() {
				line := scanner.Text()
				switch {
				case f.w:
					result[i].num += len(strings.Fields(line))
				case f.l:
					result[i].num++
				case f.m:
					result[i].num += len(line)
				}
			}

		}()
	}

	wg.Wait()

	for _, data := range result{
		fmt.Printf("%d\t%s\n", data.num, data.filename)
	}
}
