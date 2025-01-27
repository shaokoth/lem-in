package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Not enough arguments")
	}
	filename := os.Args[1]
	fmt.Println(filename)
}
