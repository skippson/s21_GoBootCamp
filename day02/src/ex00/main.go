package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type flags struct {
	sl  bool
	f   bool
	d   bool
	ext bool
	ras string
}

func (f *flags) parse() *flags {
	flag.BoolVar(&f.sl, "sl", false, "menu")
	flag.BoolVar(&f.f, "f", false, "menu")
	flag.BoolVar(&f.d, "d", false, "menu")
	flag.BoolVar(&f.ext, "ext", false, "menu")
	flag.Parse()

	if f.ext && !f.f {
		log.Fatal("-ext without -f")
	}

	if !f.d && !f.sl && !f.f{
		f.f = true
		f.d = true
		f.sl = true
	}

	return f
}

func (f *flags) getPath() string {
	arg := os.Args

	if len(arg) == 1{
		return arg[0]
	}

	for i, path := range arg{

		if path == "-ext"{
			f.ras = arg[i+1]
		}

		if path[0] == '-' || path[0] == '.'{
			continue
		}

		if path[0] != '-' && path[0] != '.' && arg[i - 1] != "-ext"{
			return path
		}

	}

	return "error path"
}

func (f flags) walkFunc(path string, info os.FileInfo, err error) error{
	_, rights := os.Stat(path)
	if rights != nil{
		if os.IsPermission(err){
			return nil
		}
	}

	linkInfo, _ := os.Lstat(path)

	switch {
	case err != nil:
		return nil
	case f.sl && linkInfo.Mode()&os.ModeSymlink == os.ModeSymlink:
		targetPath, linkErr := os.Readlink(path)
		if linkErr != nil{
			return nil
		}
		fullTargetPath := filepath.Join(filepath.Dir(path), targetPath)
		_, fullLinkErr := os.Stat(fullTargetPath)
		if fullLinkErr != nil{
			fmt.Printf("%s -> [broken]\n", path)
		} else{
			orig := path[:len(path) - len(info.Name())] + targetPath
			fmt.Printf("%s -> %s\n", path, orig)
		}
	case f.f && !info.IsDir() && linkInfo.Mode()&os.ModeSymlink != os.ModeSymlink:
		if f.ext {
			if strings.HasSuffix(path, f.ras){
				fmt.Println(path)
			}
		} else {
			fmt.Println(path)
		}
	case f.d && info.IsDir():
		fmt.Println(path)
	}

	return nil
}

func main() {
	flags := flags{}
	flags.parse()
	path := flags.getPath()

	if path == "error path" {
		log.Fatal("Invalid path")
	}

	filepath.Walk(path, flags.walkFunc)
}
