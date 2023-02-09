package main

import "fmt"

func main() {

	var keys []string

	dictionary := map[string]string{"up": "above"} // key : values
	keys = append(keys, "up")
	dictionary["1"] = "anything1"
	keys = append(keys, "1")
	dictionary["2"] = "anything2"
	keys = append(keys, "2")
	dictionary["3"] = "anything3"
	keys = append(keys, "3")
	dictionary["4"] = "anything4"
	keys = append(keys, "4")
	dictionary["5"] = "anything5"
	keys = append(keys, "5")
	fmt.Println(dictionary)

	////when ranging over maps values would not be returned in the order
	//for k, v := range dictionary {
	//	fmt.Println(k, v)
	//}

	for _, k := range keys {
		fmt.Println(dictionary[k])
	}
}
