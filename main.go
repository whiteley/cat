package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	var (
		countFlag = flag.Bool("count", false, "print number of bytes read")
	)
	flag.Parse()

	file, err := os.Open(flag.Args()[0])
	if err != nil {
		log.Fatal(err)
	}

	if count, err := io.Copy(os.Stdout, file); err != nil {
		log.Fatal(err)
	} else if *countFlag == true {
		fmt.Println(count, "bytes read")
	}
}
