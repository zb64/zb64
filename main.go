package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

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
	usePipe    bool
	rawContent string
	fileName   string
)

func throwError(m string, c int) {
	fmt.Fprintln(os.Stderr, m)
	os.Exit(c)
}

func main() {
	flag.BoolVar(&isEncode, "encode", disabled, "encode the input")
	flag.BoolVar(&isEncode, "e", disabled, "encode the input"+flagPostfix)
	flag.BoolVar(&isDecode, "decode", disabled, "decode the input")
	flag.BoolVar(&isDecode, "d", disabled, "decode the input"+flagPostfix)
	flag.BoolVar(&usePipe, "pipe", disabled, "input from stdin pipe")
	flag.BoolVar(&usePipe, "p", disabled, "input from stdin pipe"+flagPostfix)
	flag.StringVar(&rawContent, "raw", empty, "raw input content (override 'pipe')")
	flag.StringVar(&rawContent, "r", empty, "raw input content (override 'pipe')"+flagPostfix)
	flag.StringVar(&fileName, "file", empty, "input file name (override 'raw')")
	flag.StringVar(&fileName, "f", empty, "input file name (override 'raw')"+flagPostfix)
	flag.Parse()

	// source
	var content string
	if ystring.IsNotBlank(fileName) {
		if b, err := ioutil.ReadFile(fileName); err != nil {
			throwError(fmt.Sprintf("fail to read file %q: %v", fileName, err), 2)
		} else {
			content = string(b)
		}
	} else if ystring.IsNotBlank(rawContent) {
		content = rawContent
	} else if usePipe {
		fi, err := os.Stdin.Stat()
		if err != nil {
			throwError(err.Error(), 3)
		}

		if fi.Mode()&os.ModeNamedPipe != 0 {
			if data, err := ioutil.ReadAll(os.Stdin); err != nil {
				throwError(err.Error(), 4)
			} else {
				content = string(data)
			}
		} else {
			throwError("got no stdin pipeline data", 5)
		}
	} else {
		flag.Usage()
		throwError("got no input arguments, specify '-file' or '-raw'", 1)
	}

	// encode or decode
	switch {
	case !containsOneTrue(isEncode, isDecode):
		flag.Usage()
		throwError("invalid method argument, specify '-encode' or '-decode'", 10)
	case isEncode:
		if encodeResult, err := lib.Deflate.Encode([]byte(content)); err != nil {
			throwError(fmt.Sprintf("fail to encode: %v", err), 11)
		} else {
			fmt.Print(encodeResult)
		}
	case isDecode:
		if decodeResult, err := lib.Deflate.Decode(strings.TrimSpace(content)); err != nil {
			throwError(fmt.Sprintf("fail to decode: %v", err), 12)
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
