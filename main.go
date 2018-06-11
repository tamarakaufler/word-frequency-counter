// CLI tool for processing a file content using a worker pool
// for concurrency. The output gives statistics regarding occurrence of each number.
package main

import (
	"flag"
	"fmt"

	"github.com/tamarakaufler/word-frequency-counter/processor"
)

var (
	workers int
	file    string
)

func init() {
	flag.IntVar(&workers, "workers", 2, "workers flag sets the number of workers in the pool")
	flag.StringVar(&file, "file", "./test.txt", "Path to the file to be processed. Path can be relative.")
}

func main() {
	flag.Parse()

	// set up the processor
	p := &processor.Processor{
		File:    file,
		Workers: workers,
	}

	// run the task
	err := p.Run()
	if err != nil {
		fmt.Printf("ERROR: %v\n", err)
		return
	}
}
