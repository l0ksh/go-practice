// In Go, an array is a numbered sequence of elements of a specific length. 

package main

import "fmt"

func main() {
	
	var a [5]int
	fmt.Println("emp:",a)
	
	a[4] = 100
	fmt.Println("set:",a)
	fmt.Println("get:",a[4])

	fmt.Println("len:",len(a))
	
	b := [5]int{1,2,3,4,5}	//Use this syntax to declare and initialize an array in one line.
	fmt.Println("dcl:",b)
	
	var twoD [2][3]int		//Array types are one-dimensional, but you can compose types to build multi-dimensional data structures.
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD [i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
