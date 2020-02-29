package main

import (
	"github.com/burhon94/zip-archivator/pkg"
	"io/ioutil"
	"log"
	"sync"
	"testing"
)

func BenchmarkCreateZipFile(b *testing.B) {
	wg := sync.WaitGroup{}
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Printf("can't read dir: %v", err)
	}
	for _, file := range files {
		log.Printf("try create zip file: %s", file.Name())
		wg.Add(1)
		go func(wg *sync.WaitGroup,file string) {
			err := pkg.CreateZipFile(file, "./zipFiles/")
			if err != nil {
				log.Printf("can't zipped file: %v", err)
				return
			}
		}(&wg, file.Name())
		wg.Done()
	}
	wg.Wait()
}
