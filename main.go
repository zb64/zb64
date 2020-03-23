package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/1set/gut/ystring"
)

var (
	method     string
	fileName   string
	rawContent string
)

const (
	EMPTY         = ""
	METHOD_ENCODE = "encode"
	METHOD_DECODE = "decode"
	FLAG_POSTFIX  = " (shorthand)"
)

func main() {
	flag.StringVar(&method, "method", METHOD_ENCODE, "encode or decode the input")
	flag.StringVar(&method, "m", METHOD_ENCODE, "encode or decode the input"+FLAG_POSTFIX)
	flag.StringVar(&fileName, "file", EMPTY, "input file name (override 'raw')")
	flag.StringVar(&fileName, "f", EMPTY, "input file name (override 'raw')"+FLAG_POSTFIX)
	flag.StringVar(&rawContent, "raw", EMPTY, "raw input content")
	flag.StringVar(&rawContent, "r", EMPTY, "raw input content"+FLAG_POSTFIX)
	flag.Parse()

	if method != METHOD_ENCODE && method != METHOD_DECODE {
		fmt.Printf("invalid method: %s\n", method)
		flag.Usage()
		os.Exit(1)
	}

	if ystring.IsNotBlank(fileName) {
		if b, err := ioutil.ReadFile(fileName); err != nil {
			fmt.Printf("fail to read file %q: %v\n", fileName, err)
			os.Exit(3)
		} else {
			rawContent = string(b)
		}
	} else if ystring.IsBlank(rawContent) {
		fmt.Println("got no input arguments, specify 'file' or 'raw'")
		flag.Usage()
		os.Exit(2)
	}

	fmt.Printf("content: {%s}\n", rawContent)
}
