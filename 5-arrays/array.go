package main

import "fmt"

// arrays size is fixed
func main() {
	var a [5]int // 5 zeros would be the default value here
	a[0] = 100   // updating the value at the index
	a[3] = 200

	var b = [5]int{10, 20, 30} //[10 20 30 0 0]
	fmt.Printf("%#v\n", b)

	c := [...]int{1, 2, 3, 4, 5} // use len according to the number of elems provided

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(len(c))

	for i, v := range b { // i = index and v = values
		fmt.Println(i, v)
	}

	for _, v := range b { // _ = ignore values
		fmt.Println(v)
	}

	for i := range b { // i = index
		fmt.Println(i)
	}
}
