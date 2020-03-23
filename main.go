package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	host, _ := os.Hostname()
	pwd, _ := os.Getwd()
	fmt.Printf("Host: %s\nPath: %s\nTime: %s\n", host, pwd, time.Now().Format("2006-01-02T15:04:05-0700"))
}
