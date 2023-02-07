package main

import (
	"dxc-learn-go/calc"
	"dxc-learn-go/calc/data"
)

func main() {
	calc.Add()
	//fmt.Println(calc.Sum)

	data.DBConn = "connect:to:mysql"
	data.GetData()
	
	//rand.Int()
	//math.Acosh() // uncomment these to see import paths

}
