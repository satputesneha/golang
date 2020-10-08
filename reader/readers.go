package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {

	tki := strings.NewReader("reads several characters containg many characters")
	tb := make([]byte, 20)
	fmt.Println(tb)
	fmt.Println(tki.Read(tb))
	fmt.Println(tb)
	for {
		n, err := tki.Read(tb)
		fmt.Printf("n = %v err = %v tb = %v\n", n, err, tb)
		fmt.Printf("tb[:n] = %q\n", tb[:n])
		if err == io.EOF {
			break
		}

	}
}
