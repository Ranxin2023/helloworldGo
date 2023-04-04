package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")
	var x = 3
	fmt.Println(x)
	var arr = [3]int{2, 3, 4}
	fmt.Println(arr)
	for i := 0; i < 3; i++ {
		fmt.Print(arr[i], " ")
	}
}
