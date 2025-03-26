package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Kuch toh daal")
	}

	const (
		lhttp = "http://"
		uhttp = "HTTP://"
		nlink = len(lhttp)
		mask  = '*'
	)

	var (
		text = os.Args[1]
		size = len(text)
		buf  = []byte(text)
		in   bool
	)

	for i := 0; i < size; i++ {

		if len(text[i:]) >= nlink && (text[i:i+nlink] == lhttp) {
			in = true
			i += nlink
		}

		c := text[i]

		switch c {
		case ' ', '\t', '\n':
			in = false
		}

		if in {
			buf[i] = mask
		}

		buf = append(buf, c)
	}

	// print out the buffer as text (string)
	fmt.Println(string(buf))

}
