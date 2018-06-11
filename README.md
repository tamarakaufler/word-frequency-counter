# word-frequency-counter
Word Frequency Counter is a CLI tool for finding
word occurence in provided text.

 - Go
 - concurrency
 - worker pool

## Synopsis
Word Frequency Counter is implemented in Go, taking advantage
of concurrency through the use of goroutines and channels. 

Text is split into segments (sentences) which are sent to a jobs channel.
A pool of workers (gouroutines) is processing tasks (jobs) sent to the jobs
channel until all is done.

## Usage
go run main.com -file=$file_path -workers=$no_of_gouroutines

where
  - file ...... absolute or relative path of the file to be processed
  - workers ... number of workers processing the file content
