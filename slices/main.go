package main

import (
	"fmt"
)

func main() {

	//empty a slice

	s := []string{"first", "second", "third"}
	fmt.Println(s, len(s))

	clear(s)
	fmt.Println(s, len(s))

	//declaring the slice in different way

	var data []int
	fmt.Println(data, len(data), cap(data))

	var xx = []int{}
	fmt.Println(xx, len(xx), cap(xx))

	//copying slices
	x := make([]string, 0, 5)
	x = append(x, "a", "b", "c", "d")
	y := make([]string, 4)
	num := copy(y, x)
	fmt.Println(y, num)
	z := make([]string, 2)
	num1 := copy(z, y)
	fmt.Println(z, num1)

	//converting Arrays to slices

	xArray := [4]int{5, 6, 7, 8}
	xSlice := xArray[:]

	fmt.Println(xSlice)
	xArray[3] = 100 //changing the array will chnage the slice
	fmt.Println(xSlice)

	//converting the slices to arrary
	xSlice1 := []int{1, 3, 4}
	zArray1 := [3]int(xSlice1)
	fmt.Println(zArray1)
	xSlice1 = append(xSlice1, 5)
	fmt.Println(zArray1)
	fmt.Println("xSlice1: ", xSlice1)

}
