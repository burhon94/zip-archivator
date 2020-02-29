package main

import (
	"github.com/burhon94/zip-archivator/pkg"
	"io/ioutil"
	"log"
	"testing"
)

func BenchmarkCreateZipFile(b *testing.B) {
	filesDir, err := ioutil.ReadDir(".")
	if err != nil {
		log.Printf("can't read dir: %v", err)
	}
	for _, file := range filesDir {
		_ = pkg.CreateZipFile(file.Name(), "zipFiles/")
	}
}
