package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	greet()
	fmt.Println("end of main")
}

func greet() {
	// os.Args is a string type so no matter what we pass that would be always be of string type so conversion is imp
	details := os.Args[1:]

	if len(details) < 3 {
		log.Println("please provide name, age and marks")
		//log.Fataln("") // stop the program
		//os.Exit(1)

		return // stops the exec of the current func // note :- it doesn't stop the program
	}

	fmt.Println(details)
	//var err error // default val is nil // it indicates no error

	name := details[0]
	ageString := details[1]
	marksString := details[2]

	// john , abc , xyz
	age, err := strconv.Atoi(ageString)
	if err != nil {
		log.Println(err)
		fmt.Println(age) // age is set to its default when err occurs
		return           // stops the exec of the current func
	}

	marks, err := strconv.Atoi(marksString)

	if err != nil {
		log.Println(err)
		return // stops the exec of the current func
	}

	fmt.Println(name, age, marks)

	//2nd // not a good approach
	//if err != nil {
	//	log.Println(err)
	//} else {
	//	fmt.Println(name)
	//	fmt.Println(age)
	//	/* err = sendPipeline(name,age)
	//	 if err != nil {
	//		log error
	//	}else {
	//
	//	}
	//	*/
	//}

}
