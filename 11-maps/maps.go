package main

import "fmt"

func main() {

	//var dictionary map[string]string // default value is nil
	dictionary := make(map[string]string)

	dictionary["up"] = "down"
	dictionary["down"] = "below"
	fmt.Println(dictionary)

	delete(dictionary, "up")
	for k, v := range dictionary {
		fmt.Println(k, v)
	}

}
