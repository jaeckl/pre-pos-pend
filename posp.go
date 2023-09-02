//go:build posp

package main

import (
	"fmt"
	"io"
	"github.com/jaeckl/pre-pos-pend/internal"
)

func main() {
	internal.Run(func(writer io.Writer, arg string, sep string, text string) {
		_, err := writer.Write([]byte(fmt.Sprintf("%s%s%s\n", text, sep, arg)))
		if err != nil {
			panic(err)
		}
	})
}
