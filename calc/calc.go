// Package calc is library package
package calc

import "fmt"

var sum int // global var // local to the package

// Add func is exported as first letter is uppercase
func Add() {

	i, b := 10, 20
	sum = i + b
	fmt.Println(sum)
	sub()
	//fmt.Sprintf() // peek into this for how an unexported func would be used
}

func ABC() {

}
