package main

import "fmt"

func main() {

	//how we can define the variables
	//Type 1 -> we need to define the type or initialize it. One thing is required to define
	var heyGo string = "Go on a hike"

	//Type 2 -> we need to assign the value to the variable
	var withWhome = "Python"

	//Type 3 -> we need to assign the value to the variable and this type of variable is only used inside the function
	howManyPeople := 2

	//Type 4 -> Define multiple variables at a time
	var a, b = "hey", "you"
	c, d := "hellp", "c"

	//Type 5 -> Define a variable in the block form
	var (
		e int  = 1
		f bool = true
	)

	// How we can print things in go
	fmt.Println("Hello, World!")
	fmt.Println(heyGo + "with " + withWhome)
	fmt.Println("I'm going to hike with", howManyPeople, "people")
	fmt.Println("Multiple variable type 1", a, "and", b)
	fmt.Println("Multiple variable type 2", c, "and", d)
	fmt.Println("Multiple variable type 3", e, "and", f)

}
