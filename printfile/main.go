package main

import (
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open(os.Args[1])

	if err != nil {
		log.Fatalln(err)
	}

	io.Copy(os.Stdout, file)
}
