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
	// third exercise, slices of slices
	r2d2 := [][]string{
		[]string{"a", "b", "c", "d"},
		[]string{"e", "f", "g", "h"},
		[]string{"i", "j", "k", "l"},
		[]string{"m", "n", "o", "p"},
	}
	fmt.Println("Starting slice of slices:")
	fmt.Println(r2d2, len(r2d2), cap(r2d2))
	for x := 0; x < len(r2d2); x++ {
		fmt.Println(r2d2[x])
	}
	// // rotating our cube by 90 degrees
	// r2d2_rotate := make([][]string, 4)
	// // copy(r2d2_rotate, r2d2)
	// // fmt.Println(r2d2_rotate)
	// for x := 0; x < len(r2d2); x++ {
	// 	for y := 0; y < len(r2d2[x]); y++ {
	// 		r2d2_rotate[x][y] = r2d2[]
	// 	}
	// }
}
