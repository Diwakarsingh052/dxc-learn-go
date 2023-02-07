package main

import (
	"fmt"
	"os"
)

func main() {

	args := os.Args[1:] // [progName,arg ...] // from 1st index till the end of list

	fmt.Println(args[0])
}
