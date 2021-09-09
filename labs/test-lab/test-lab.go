package main

import (
	"fmt";
	"os"
)

func main() {

	n := ""
	for i := 1; i < len(os.Args); i++ {
		c := os.Args[i]
		c += " "
		n += c 
	}
	fmt.Printf("%v, Welcome to the jungle\n", n)
}