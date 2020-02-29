package main

import (
	"fmt"
	"github.com/burhon94/zip-archivator/pkg"
	"os"
)

var zipDir = "./zipFiles/"

func main() {
	files := os.Args[1:]
	fmt.Printf("this files will zip: %s\n", files)
	for _, file := range files {
		fmt.Printf("check exist file: %s\n", file)
		_, err := os.Stat(file)
		if err != nil {
			fmt.Printf("can't find file: %s\n", file)
			continue
		}
		fmt.Printf("try create zip file: %s\n", file)
		err = pkg.CreateZipFile(file, zipDir)
		if err != nil {
			fmt.Printf("can't zipped file: %v\n", err)
			return
		}
		fmt.Printf("file %s zipped success\n", zipDir+file)
	}
}