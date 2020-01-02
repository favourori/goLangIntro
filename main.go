package main

import (
	"fmt"
)

func main() {
	//fmt.Println("Hello Go!")

	//Datatypes
	var name string = "Favour Ori"
	var age int = 23
	const isloggedIn = false
	fmt.Println(name, age, isloggedIn)

	//print data type
	fmt.Printf("%T\n", age)
}
