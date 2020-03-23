package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/1set/gut/ystring"
)

const (
	EMPTY         = ""
	METHOD_ENCODE = "encode"
	METHOD_DECODE = "decode"
	FLAG_POSTFIX  = " (shorthand)"
)

var (
	method     string
	fileName   string
	rawContent string
)

func main() {
	flag.StringVar(&method, "method", METHOD_ENCODE, "encode or decode the input")
	flag.StringVar(&method, "m", METHOD_ENCODE, "encode or decode the input"+FLAG_POSTFIX)
	flag.StringVar(&fileName, "file", EMPTY, "input file name (override 'raw')")
	flag.StringVar(&fileName, "f", EMPTY, "input file name (override 'raw')"+FLAG_POSTFIX)
	flag.StringVar(&rawContent, "raw", EMPTY, "raw input content")
	flag.StringVar(&rawContent, "r", EMPTY, "raw input content"+FLAG_POSTFIX)
	flag.Parse()

	var content string
	if ystring.IsNotBlank(fileName) {
		if b, err := ioutil.ReadFile(fileName); err != nil {
			fmt.Printf("fail to read file %q: %v\n", fileName, err)
			os.Exit(3)
		} else {
			content = string(b)
		}
	} else if ystring.IsNotBlank(rawContent) {
		content = rawContent
	} else {
		fmt.Println("got no input arguments, specify 'file' or 'raw'")
		flag.Usage()
		os.Exit(2)
	}

	fmt.Printf("content: {%s}\n", rawContent)
	switch method {
	case METHOD_ENCODE:
		if encodeResult, err := base64Deflate.Encode([]byte(content)); err != nil {
			fmt.Printf("fail to encode: %v\n", err)
			os.Exit(4)
		} else {
			fmt.Println(encodeResult)
		}
	case METHOD_DECODE:
		if decodeResult, err := base64Deflate.Decode(content); err != nil {
			fmt.Printf("fail to decode: %v\n", err)
			os.Exit(5)
		} else {
			fmt.Println(string(decodeResult))
		}
	default:
		fmt.Printf("invalid method: %s\n", method)
		flag.Usage()
		os.Exit(1)
	}
}
