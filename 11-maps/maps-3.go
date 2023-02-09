package main

import (
	"fmt"
)

type user struct {
	name string
}

func main() {
	u := make(map[int]*user) //k = int , v = user
	u[101] = &user{name: "john"}

	v, ok := u[102]

	//avoid this
	//if v == nil {
	//	fmt.Println("not there ")
	//	return
	//}

	if !ok {
		return
	}

	fmt.Printf("%+v", v)
}
