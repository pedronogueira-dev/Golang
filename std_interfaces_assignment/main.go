package main

import (
	"io"
	"os"
)

func main() {
	fname := os.Args[1]
	if fname == "" {
		fname = "myfile.txt"
	}
	var file *os.File
	file, err := os.Open(fname)

	if err != nil {
		file, _ = os.Open("myfile.txt")
	}

	io.Copy(os.Stdout, file)
}
