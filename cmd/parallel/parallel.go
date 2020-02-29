package main

import (
	"fmt"
	"github.com/burhon94/zip-archivator/pkg"
	"os"
	"sync"
)

var zipDir = "./zipFiles/"

func main() {
	wg := sync.WaitGroup{}
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
		wg.Add(1)
		go func(wg *sync.WaitGroup,file string) {
			defer func() {
				wg.Done()
			}()
			err := pkg.CreateZipFile(file, zipDir)
			if err != nil {
				fmt.Printf("can't zipped file: %v\n", err)
				return
			}
			fmt.Printf("file %s zipped success\n", zipDir+file)
		}(&wg, file)
	}
	wg.Wait()
}
