package main

import (
	"fmt"
	"os"
	"strings"
)

type void struct {}
var member void

func main() {

	pathstring := os.Getenv("PATH")
	paths := strings.Split(pathstring, ":")
	x := make(map[string]void)

	for _, path := range paths {
		execs, _ := os.ReadDir(path)
		for _, executable := range execs {
			x[executable.Name()] = member
		}
	}

	for k := range x {
		fmt.Println(k)
	}
}
