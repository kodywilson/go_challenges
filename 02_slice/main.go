package main

import "fmt"

func main() {

	var slicey []int
	fmt.Println("This is what we start with:")
	fmt.Println(slicey, len(slicey), cap(slicey))
	fmt.Println("---------------------------")
	for i := 1; i < 11; i++ {
		slicey = append(slicey, i)
		fmt.Println("val: ", i, "size of slice: ", len(slicey), "capacity of slice: ", cap(slicey))
	}
}
