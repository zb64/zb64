package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/1set/gut/ystring"
	"github.com/zb64/lib"
)

const (
	disabled    = false
	empty       = ""
	flagPostfix = " (shorthand)"
)

var (
	isEncode   bool
	isDecode   bool
	fileName   string
	rawContent string
)

func main() {
	flag.BoolVar(&isEncode, "encode", disabled, "encode the input")
	flag.BoolVar(&isEncode, "e", disabled, "encode the input"+flagPostfix)
	flag.BoolVar(&isDecode, "decode", disabled, "decode the input")
	flag.BoolVar(&isDecode, "d", disabled, "decode the input"+flagPostfix)
	flag.StringVar(&fileName, "file", empty, "input file name (override 'raw')")
	flag.StringVar(&fileName, "f", empty, "input file name (override 'raw')"+flagPostfix)
	flag.StringVar(&rawContent, "raw", empty, "raw input content")
	flag.StringVar(&rawContent, "r", empty, "raw input content"+flagPostfix)
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
	case !containsOneTrue(isEncode, isDecode):
		fmt.Println("invalid method argument, specify '-encode' or '-decode'")
		flag.Usage()
		os.Exit(3)
	case isEncode:
		if encodeResult, err := lib.Deflate.Encode([]byte(content)); err != nil {
			fmt.Printf("fail to encode: %v\n", err)
			os.Exit(4)
		} else {
			fmt.Print(encodeResult)
		}
	case isDecode:
		if decodeResult, err := lib.Deflate.Decode(content); err != nil {
			fmt.Printf("fail to decode: %v\n", err)
			os.Exit(5)
		} else {
			fmt.Print(string(decodeResult))
		}
	}
}

func containsOneTrue(vals ...bool) bool {
	count := 0
	for _, v := range vals {
		if v {
			count++
		}
	}
	return count == 1
}
