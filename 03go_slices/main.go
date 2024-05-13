package main

import "fmt"

func main() {
	sliceOne := []int{1, 2, 3, 4}
	fmt.Println(sliceOne)
	fmt.Println(len(sliceOne))
	fmt.Println(cap(sliceOne))

	// From array to slice

	arr1 := [6]int{10, 11, 12, 13, 14,15}
  	myslice := arr1[0:4]
	fmt.Println(myslice)
	fmt.Println(len(myslice))
	fmt.Println(cap(myslice))


	sliceFromMake := make([]int,5,6)

	fmt.Println("sliceFromMake",sliceFromMake)
}