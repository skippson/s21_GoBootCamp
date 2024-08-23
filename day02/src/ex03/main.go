package main

import (
	"archive/tar"
	"compress/gzip"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

func main() {
	files, directory := getFiles()


	var wg sync.WaitGroup
	wg.Add(len(files))

	for _, file := range files {
		arch(file, directory, &wg)
	}

	wg.Wait()

}

func arch(filename string, directory string, wg *sync.WaitGroup) {
	defer wg.Done()

	name := fmt.Sprintf("%s_%d",trimName(filename, directory), time.Now().Unix())
	tarFile(filename, name)

}

func tarFile(filename string, nameArch string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	tarFile, err := os.Create(nameArch + ".tar")
	if err != nil {
		log.Fatal(err)
	}
	defer tarFile.Close()

	tarWriter := tar.NewWriter(tarFile)
	info, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}

	header, err := tar.FileInfoHeader(info, file.Name())
	if err != nil {
		log.Fatal(err)
	}

	err = tarWriter.WriteHeader(header)
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.Copy(tarWriter, file)
	if err != nil {
		log.Fatal(err)
	}

	err = tarWriter.Close()
	if err != nil {
		log.Fatal(err)
	}

	gzTar(tarFile, nameArch)
}

func gzTar(tarFile *os.File, nameArch string) {
	gzipFile, err := os.Create(nameArch + ".tar.gz")
	if err != nil {
		log.Fatal(err)
	}
	defer gzipFile.Close()

	gzipWriter := gzip.NewWriter(gzipFile)
	defer gzipWriter.Close()

	_, err = io.Copy(gzipWriter, tarFile)
	if err != nil {
		log.Fatal(err)
	}

	err = os.Remove(tarFile.Name())
	if err != nil {
		log.Fatal(err)
	}
}

func trimName(filename string, directory string) string {
	idx := strings.LastIndex(filename, ".")
	if idx > 0{
		filename = filename[0:idx]
	}

	if len(directory) > 0{
		for counts := count(filename); counts > 0; counts-- {
			_, filename, _ = strings.Cut(filename, "/")
		}
	
		if strings.HasSuffix(directory, "/") {
			filename = directory + filename
		} else {
			filename = directory + "/" + filename
		}
	}

	return filename
}

func count(s string) int {
	ans := 0
	for _, data := range s {
		if data == '/' {
			ans++
		}
	}

	return ans
}

func mkdir(dir string) string {
	_, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			err := os.Mkdir(dir, 0755)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	if !strings.HasSuffix(dir, "/") {
		dir += "/"
	}

	return dir
}

func getFiles() ([]string, string) {
	flagDir := flag.Bool("a", false, "Name of directory")
	flag.Parse()
	files := os.Args[1:]

	if *flagDir {
		directory := mkdir(files[1])
		mkdir(directory)
		files = files[2:]

		return files, directory
	}

	return files, ""
}
