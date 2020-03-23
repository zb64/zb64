package main

import (
	"flag"
	"fmt"
	"os"
	"time"

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

	fmt.Printf("m: %s\n", method)
	fmt.Printf("f: %s\n", fileName)
	fmt.Printf("r: %s\n", rawContent)

	if method != METHOD_ENCODE && method != METHOD_DECODE {
		fmt.Printf("invalid method: %s\n", method)
		flag.Usage()
		os.Exit(1)
	}

	if ystring.IsNotBlank(fileName) {
		fmt.Println("read", fileName)
	} else if ystring.IsNotBlank(rawContent) {
		fmt.Println("got", rawContent)
	} else {
		fmt.Println("got no input arguments, specify 'file' or 'raw'")
		flag.Usage()
		os.Exit(2)
	}

	host, _ := os.Hostname()
	pwd, _ := os.Getwd()
	fmt.Printf("Host: %s\nPath: %s\nTime: %s\n", host, pwd, time.Now().Format("2006-01-02T15:04:05-0700"))
}
