package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"l2.9/pkg/strunpack"
)

func main() {
	path := flag.String("file", "examples.txt", "path to examples")
	file, err := os.ReadFile(*path)
	if err != nil {
		fmt.Println("smth wrong")
	}
	s := strings.SplitSeq(string(file), "\n")
	for v := range s {
		res, err := strunpack.StrUnpack(v)
		if err != nil {
			fmt.Printf("Error with string: %s, \t%v\n", v, err.Error())
			continue
		}
		fmt.Printf("row string: %s, \tunpacked string: %s\n", v, res)
	}
}
