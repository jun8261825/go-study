// myMath.go

package myMath

import "fmt"

func Add(x, y int) int {
	return x + y
	fmt.Println(x + y)
}

func Sub(x, y int) int {
	return x - y
}

func Max(x, y int) int {
	var result int

	if x > y {
		result = x
	} else {
		result = y
	}
	return result
}
