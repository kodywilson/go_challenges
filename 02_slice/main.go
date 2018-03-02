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

	// slice of spanish numbers
	pie := []string{"uno", "dos", "tres", "cuatro", "cinco"}
	fmt.Println(pie, len(pie), cap(pie))
	for x := 0; x < len(pie); x++ {
		fmt.Println(pie[x])
	}
	// same numbers counting down
	for x := (len(pie) - 1); x >= 0; x-- {
		fmt.Println(pie[x])
	}
	// could also create a new slice
	var pizza []string
	for x := (len(pie) - 1); x >= 0; x-- {
		pizza = append(pizza, pie[x])
	}
	fmt.Println(pizza)
}
