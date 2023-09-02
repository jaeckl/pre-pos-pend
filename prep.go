//go:build prep

package main

import (
	"fmt"
	"io"
	"github.com/jaeckl/pre-pos-pend/internal"
)

func main() {
	internal.Run(func(writer io.Writer, arg string, sep string, text string) {
		_, err := writer.Write([]byte(fmt.Sprintf("%s%s%s\n", arg, sep, text)))
		if err != nil {
			panic(err)
		}
	})
}
