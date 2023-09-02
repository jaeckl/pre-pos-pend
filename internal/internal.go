package internal

import (
	"bufio"
	"flag"
	"io"
	"os"
)

const NO_VALUE = ""
const FLAG_SEPERATOR = "s"
const FLAG_INPUTPATH = "i"
const FLAG_OUTPUTPATH = "o"

func isSetByUser(flagName string) bool {
	var visited bool = false
	flag.Visit(func(f *flag.Flag) {
		if f.Name == flagName {
			visited = true
		}
	})
	return visited
}

func Run(processor func(writer io.Writer, arg string, sep string, text string)) {
	var flagSeperator = flag.String(FLAG_SEPERATOR, NO_VALUE, "seperator after the argument")
	var flagInputFile = flag.String(FLAG_INPUTPATH, NO_VALUE, "path to input file")
	var flagOutputFile = flag.String(FLAG_OUTPUTPATH, NO_VALUE, "path to ouput file")

	var reader io.Reader = os.Stdin
	var writer io.Writer = os.Stdout
	var inFile *os.File
	var ouFile *os.File

	flag.Parse()
	args := flag.Args()

	if len(args) > 1 || len(args) < 1 {
		panic("Expecting one parameter")
	}

	if isSetByUser(FLAG_INPUTPATH) {
		var err error
		inFile, err = os.Open(*flagInputFile)
		if err != nil {
			panic(err)
		}
		reader = bufio.NewReader(inFile)
	}
	if isSetByUser(FLAG_OUTPUTPATH) {
		var err error
		ouFile, err = os.Create(*flagOutputFile)
		if err != nil {
			panic(err)
		}
		writer = bufio.NewWriter(ouFile)
	}
	defer func() {
		if isSetByUser(FLAG_INPUTPATH) {
			if err := inFile.Close(); err != nil {
				panic(err)
			}
		}
		if isSetByUser(FLAG_OUTPUTPATH) {
			writer.(*bufio.Writer).Flush()
			if err := ouFile.Close(); err != nil {
				panic(err)
			}
		}
	}()

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		processor(writer, args[0], *flagSeperator, scanner.Text())
	}
}
