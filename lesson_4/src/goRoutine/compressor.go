package main

import (
	"log"
	"sync"
	"time"
)

/*
Компрессор файлов,
имена которых будут указаны в аргументах
терминала при вызове.
*/

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

func main() {
	var wg sync.WaitGroup

	now := time.Now()

	for ind, arg := range os.Args {
		fmt.Printf("%d. %s\n", ind+1, arg)
	}
	for _, fileName := range os.Args[1:] {
		wg.Add(1)
		go func(fileName string, wg *sync.WaitGroup) {
			if err := compress(fileName, wg); err != nil {
				log.Printf("Error: %s",  err)
			}
		}(fileName, &wg)

	}
	wg.Wait()
	fmt.Println(time.Since(now))
}

func compress(fileName string, wg *sync.WaitGroup) (err error) {
	file, err := os.Open(fileName)
	if err != nil {
		return
	}
	defer file.Close()

	out, err := os.Create(fileName + ".gz")
	if err != nil {
		return
	}
	defer out.Close()

	gzFile := gzip.NewWriter(out)
	if _, err = io.Copy(gzFile, file); err != nil {
		return
	}
	defer gzFile.Close()
	wg.Done()
	return
}
