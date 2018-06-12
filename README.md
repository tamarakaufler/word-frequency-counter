# word-frequency-counter
Word Frequency Counter is a CLI tool for finding
word occurences in provided text.

 - Go
 - concurrency
 - worker pool
 - reduce

## Synopsis
Word Frequency Counter is implemented in Go, taking advantage
of concurrency through the use of goroutines and channels. 

Text is split into segments (sentences) which are sent to a jobs channel.
A pool of workers (gouroutines) is processing tasks (jobs) sent to the jobs
channel until all is done.

## Usage
go run main.com -file=$file_path -workers=$no_of_gouroutines

docker run --name=wordcounter --rm -v $PWD:/data quay.io/tamarakaufler/wordcounter:v1alpha1 --file=/data/some_file.txt --workers=4

  where some_file.txt is a file in the directory from which the docker command is run.

where
  - file ...... absolute or relative path of the file to be processed
  - workers ... number of workers processing the file content
