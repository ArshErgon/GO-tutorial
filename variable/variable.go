package main

import (
	"fmt"
	"strconv"
)

var new_i int = 49

func main() {
	// there are three ways to declare the variable in GO
	
	// var i int32 = 30
	num_one := 30
	num_two := 40

	sum_ := num_one + num_two

	fmt.Printf("%v, %T\n", sum_, sum_)
	fmt.Printf("%v", new_i)

	// multi line declaring the variables

	var (
		one int = 1
		two int = 2
	)

	fmt.Println(one, two)

	// you should remember that you can't let an unused variable
	// to be stay in your program.


	// type conversion

	var i int = 42
	fmt.Println(i)

	// type conversion isn't same like python or java
	// https://pkg.go.dev/strconv?utm_source=gopls#Itoa will tell you alot about string conversion

	var i_to_string string = strconv.Itoa(i)
	fmt.Printf("%v %T\n",i_to_string, i_to_string)

	var is_it bool = true

	fmt.Println(is_it)

	// bit operators
	a := 10 // 1010
	b := 3  // 0011

	fmt.Println(a & b) // 0010
	fmt.Println(a | b) // 1011
	fmt.Println(a ^ b) // 1001
	fmt.Println(a &^ b) // 0100

	fmt.Println(a << 1)
	fmt.Println(a >> 1)




}