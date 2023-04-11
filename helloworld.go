package main

import (
	"fmt"
)

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func substring(start int, end int, s string) string {
	var ret = []byte{}
	for i := start; i < end; i++ {
		ret = append(ret, s[i])
	}
	return string(ret)
}

func array_demo() {
	var arr = [3]int{2, 3, 4}
	fmt.Println(arr)
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i], " ")
	}

	fmt.Println("The length of arr is :", len(arr))
	fmt.Println("The capacity of arr is :", cap(arr))
	var slice = arr[1:3]
	slice = append(slice, 12, 13, 14, 15, 16, 17, 18, 20, 21)
	fmt.Println("The length of arr is:", len(slice))
	fmt.Println("The capacity of arr is:", cap(slice))
	slice = arr[1:3]
	slice = append(slice, 12, 13, 14, 15, 16, 17, 18)
	fmt.Println("The length of arr is:", len(slice))
	fmt.Println("The capacity of arr is:", cap(slice))
	slice = arr[1:1]
	slice = append(slice, 12)
	fmt.Println("The length of arr is:", len(slice))
	fmt.Println("The capacity of arr is:", cap(slice))
	var arr2 = [10]int{1}
	for i := range arr2 {
		fmt.Print(arr2[i], " ")
	}
	fmt.Println()
}
func main() {
	fmt.Println("Hello World!")
	var x = 3
	fmt.Println(x)
	fmt.Println(factorial(8))
	array_demo()
	fmt.Println("substring of 'abcde' from 1 to 3 is:", substring(1, 3, "abcde"))
}
