package main

import (
	"fmt"

	"github.com/kodywilson/go_challenges/00_hello/printer"
)

func main() {
	err := printer.PrintMessage("Goooooooood Morning Vietnam!!!!")
	if err == nil {
		fmt.Println("It worked!!!")
	}
}
