package main

import (
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("The Go Programming Language")

	buf := make([]byte, 8)
	for {
		n, err := r.Read(buf)
		if n > 0 {
			println(string(buf[:n]))
		}
		if err == io.EOF {
			break
		}
	}
}
