package main

import (
	"fmt"
)

func main() {

	var s string // default value is empty string of string datatype
	var i int    // 0

	fmt.Println(s, i)
	fmt.Printf("%q\n", s)

	//var name = "Bob"
	//name = 1 // go is a statically compiled language
	//_ = name // ignore values // don't use in production code

	//age, marks := 10, "20" //shorthand
	//_, _ = age, marks

	var (
		name  string
		age   = 33
		marks int
	)

	_, _, _ = name, age, marks

	//time.Second // peek into this code to see design pattern
	//os.O_APPEND // peek into this code to see design pattern
}
