package main

import "fmt"

func main() {

	if a, b := 10, 20; a > b {
		fmt.Println(a)
	} else {
		fmt.Println(b)
	} // life of var declared inside the if block ends with if

	//fmt.Println(a)

}
