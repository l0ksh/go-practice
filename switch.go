package main

import (
	"fmt"
	"time"
)

func main() {
	
	i := 2
	fmt.Print("Write ", i ," as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2: 
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's the weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
		fmt.Println(t.Hour())
	}
/*
In Go, whatAmI := func(i interface{}) is a function declaration. The func keyword indicates that the following code is a function definition. The whatAmI is the name of the function, and the i interface{} is the parameter type of the function. The interface{} type is the empty interface, which means that it can hold any type of value.
*/

	
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Print("Don't know type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
}
