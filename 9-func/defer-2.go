package main

import (
	"log"
	"os"
)

func main() {

	file, err := os.OpenFile("test.txt", os.O_RDWR, 0755)
	if err != nil {
		log.Println(err)
		return
	}

	defer file.Close() // it guarantees that file would be closed when the function ends

	// do work with your file

	// panic happened

	//file would be closed safely

}
