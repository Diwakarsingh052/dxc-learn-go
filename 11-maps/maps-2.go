package main

import (
	"fmt"
	"log"
)

type user struct {
	name string
}

func main() {
	u := make(map[int]user) //k = int , v = user
	u[101] = user{name: "john"}

	v, ok := u[101] // ok is bool // ok = true if value is found
	if !ok {        // ok == false
		log.Println("value not found")
		return
	}
	fmt.Printf("%+v", v)
}
