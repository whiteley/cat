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
		countFlag  = flag.Bool("count", false, "print number of bytes read")
		stderrFlag = flag.Bool("stderr", false, "write to stderr instead of stdout")
	)
	flag.Parse()

	file, err := os.Open(flag.Args()[0])
	if err != nil {
		log.Fatal(err)
	}

	out := os.Stdout
	if *stderrFlag == true {
		out = os.Stderr
	}

	if count, err := io.Copy(out, file); err != nil {
		log.Fatal(err)
	} else if *countFlag == true {
		fmt.Println(count, "bytes read")
	}
}
