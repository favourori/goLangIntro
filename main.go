package main

import (
	"fmt"
	"math"

	"github.com/favourori/go_intro/mathUtils"
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

	//conditionals
	var x = 50
	var y = 10
	if x < y {
		fmt.Printf("%d is less than %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}

	//loops
	//long method
	// i := 1
	// for i <= 10 {
	// 	fmt.Println(i)
	// 	i++
	// }

	//short method
	for j := 1; j <= 10; j++ {
		fmt.Println(j)

	}

	fmt.Println(mathUtils.Sum(4, 3))

}
