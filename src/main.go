package main

import (
	"bible"
	"fmt"
	"os"
	"strings"
)

func main() {
	var args = strings.Join(os.Args[1:], " ")

	if len(args) == 0 {
		bible.Help()
		os.Exit(1)
	}

	var resp = bible.Get(args)
	fmt.Println(resp)
}
