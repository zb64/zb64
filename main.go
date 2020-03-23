package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/1set/gut/ystring"
)

const (
	DISABLE     = false
	EMPTY       = ""
	FlagPostfix = " (shorthand)"
)

var (
	isEncode   bool
	isDecode   bool
	method     string
	fileName   string
	rawContent string
)

func main() {
	flag.BoolVar(&isEncode, "encode", DISABLE, "encode the input")
	flag.BoolVar(&isEncode, "e", DISABLE, "encode the input"+FlagPostfix)
	flag.BoolVar(&isDecode, "decode", DISABLE, "decode the input")
	flag.BoolVar(&isDecode, "d", DISABLE, "decode the input"+FlagPostfix)
	flag.StringVar(&fileName, "file", EMPTY, "input file name (override 'raw')")
	flag.StringVar(&fileName, "f", EMPTY, "input file name (override 'raw')"+FlagPostfix)
	flag.StringVar(&rawContent, "raw", EMPTY, "raw input content")
	flag.StringVar(&rawContent, "r", EMPTY, "raw input content"+FlagPostfix)
	flag.Parse()

	var content string
	if ystring.IsNotBlank(fileName) {
		if b, err := ioutil.ReadFile(fileName); err != nil {
			fmt.Printf("fail to read file %q: %v\n", fileName, err)
			os.Exit(2)
		} else {
			content = string(b)
		}
	} else if ystring.IsNotBlank(rawContent) {
		content = rawContent
	} else {
		fmt.Println("got no input arguments, specify '-file' or '-raw'")
		flag.Usage()
		os.Exit(1)
	}

	switch {
	case (isEncode && isDecode) || !(isEncode || isDecode):
		fmt.Println("got no method arguments, specify '-encode' or '-decode'")
		flag.Usage()
		os.Exit(3)
	case isEncode:
		if encodeResult, err := base64Deflate.Encode([]byte(content)); err != nil {
			fmt.Printf("fail to encode: %v\n", err)
			os.Exit(4)
		} else {
			fmt.Print(encodeResult)
		}
	case isDecode:
		if decodeResult, err := base64Deflate.Decode(content); err != nil {
			fmt.Printf("fail to decode: %v\n", err)
			os.Exit(5)
		} else {
			fmt.Print(string(decodeResult))
		}
	}
}
