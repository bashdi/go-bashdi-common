package common

import (
	"fmt"
)

func PrintArray(array ...interface{}) {
	fmt.Println("array start:")
	defer fmt.Println("array end;")

	for i, v := range array {
		fmt.Printf("index:\t%d \t\tvalue:\t%v\n", i, v)
	}
}
