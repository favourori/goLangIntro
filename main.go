package main

import (
	"fmt"
	"math"

	"github.com/favourori/go_intro/stringUtils"
)

func greet(name string) string {
	return "Hello, " + name

}

func main() {
	//fmt.Println("Hello Go!")

	//Datatypes
	var name string = "Favour Ori"
	var age int = 23
	const isloggedIn = false
	fmt.Println(name, age, isloggedIn)
	fmt.Println(math.Floor(20.6))
	fmt.Println(stringUtils.Reverse("Hello"))

	//print data type
	fmt.Printf("%T\n", age)

	//functions
	fmt.Println(greet("David"))

	//Arrays
	fruits := [2]string{"Orange", "Mango"}
	fmt.Println(fruits)
}
