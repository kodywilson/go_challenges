package printer

import (
	"fmt"
)

func PrintMessage(s string) error {
	_, err := fmt.Println(s)
	return err
}
