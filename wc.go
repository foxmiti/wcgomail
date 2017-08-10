package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {

	if len(os.Args) > 2 {
		panic("wrong arguments")
	}

	path := os.Args[1]
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	N := 0
	const sizeBuf = 1024
	buf := make([]byte, sizeBuf)
	for {
		n, err := file.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Printf("%v\n", err)
				return
			}
		}
		N += bytes.Count(buf[:n], []byte("\n"))
	}
	fmt.Printf("\t%v\n", N)
}
