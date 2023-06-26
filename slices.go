package main

import "fmt"

func main() {
	
	var s[]string
	fmt.Println("uninit:",s, s == nil, len(s) == 0)

	s = make([]string, 3)
	fmt.Println("emp:",s,"len:", len(s),"cap:",cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e","f")
	
	fmt.Println("apd:",s)
	
	//Slices can also be copy’d. Here we create an empty slice c of the same length as s and copy into c from s.

	c := make([]string, len(s))
	copy(c,s)
	fmt.Println("cpy:",c)

	//Slices support a “slice” operator with the syntax slice[low:high]. For example, this gets a slice of the elements s[2], s[3], and s[4].	

	l := s[2:5]
	fmt.Println("sl1:",l)

	l = s[:5]	
	fmt.Println("sl2:",l)

	l = s[2:]
	fmt.Println("sl3:",l)
	
	//We can declare and initialize a variable for slice in a single line as well.
	t := []string{"g","h","i"}
	fmt.Println("dcl:",t)

	// Slices can be composed into multi-dimensional data structures. The length of the inner slices can vary, unlike with multi-dimensional arrays.
	twoD := make([][]int,3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
