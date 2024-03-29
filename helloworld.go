package main

import (
	"fmt"
	"strconv"
	"strings"
)

var global int

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

func substring2(start int, end int, s string) string {
	return s[start:end]
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
func switch_case(x interface{}) {

	switch v := x.(type) {
	case float32:
		fmt.Println("x is ", x)
	case int:
		fmt.Println("v is ", v)
	case mytype:
		fmt.Println("global is ", global)
	default:
		fmt.Println("default is ", 1234)
	}
}
func double_dimensional_array(height int, width int) {
	var matrix = make([][]int, height)
	fmt.Printf("new matrix of %d and %d: \n[", height, width)
	for i := range matrix {
		matrix[i] = make([]int, width)
	}
	for i, row := range matrix {
		for j, value := range row {
			if i == height-1 && j == width-1 {
				fmt.Printf("%d] \t", value)
			} else {
				fmt.Printf("%d, \t", value)
			}

		}

		fmt.Println()
	}
}

type mytype float32

func inner_func() {
	var arr = []int{0, 1, 2, 3, 4, 5, 6}
	slice := func(arr []int, i int, j int) []int {
		var res = []int{}
		for idx := i; idx < j; idx++ {
			res = append(res, arr[idx])
		}
		return res
	}
	var arr_string = fmt.Sprint(slice(arr, 3, 6))
	fmt.Println("arr is: " + arr_string)
}

func map_demo() {
	var arr_idx = make(map[int][]int)
	var arr = []int{10, 8, 3, 5, 2, 11, 10, 3, 2, 2, 6}
	for i, num := range arr {
		arr_idx[num] = append(arr_idx[num], i)
	}
	for key, idxs := range arr_idx {
		fmt.Printf("element:%d: ", key)
		for _, idx := range idxs {
			fmt.Printf("%d ", idx)
		}
		fmt.Println()
	}
}
func set_demo() {
	var num_set = make(map[int]bool)
	num_set[1] = true
	num_set[7] = true
	num_set[5] = true
	num_set[12] = true
	num_set[9] = true
	num_set[5] = true
	fmt.Println("whether 5 is in num_set:", num_set[5])
	fmt.Println("whether 10 is in num_set:", num_set[10])
}
func inner_recur_demo() {
	var inner_factorial func(int) int
	inner_factorial = func(n int) int {
		if n == 0 {
			return 1
		}
		return n * inner_factorial(n-1)
	}
	fmt.Println("inner factorial 8 is", inner_factorial(8))
}
func string_convert_demo() {
	num := 123
	num_str := strconv.Itoa(num)
	fmt.Printf("123 convert to string is:%s\n", num_str)
	num_back, _ := strconv.Atoi(num_str)
	fmt.Printf("'123' convert to int is:%d\n", num_back)

}
func string_builder_demo() {
	var sentence strings.Builder
	var words = []string{"hello", "world", "this", "is", "Jack"}
	for _, word := range words {
		if sentence.Len() > 0 {
			sentence.WriteString(" ")
		}
		sentence.WriteString(word)
	}
	fmt.Printf("The sentence is %s", sentence.String())
}

func oop_demo() {
	s1 := StudentConstructor("john", "1234", 20)
	s1.Show()
	s1.EnrollCourses("Physics")
	s1.EnrollCourses("Chemistry")
	s1.EnrollCourses("Data Structure")
	s1.PrintCourses()
	s1.changeName("tom")
	s1.changeAge(21)
	s1.Show()
}

func main() {
	fmt.Println("Hello World!")
	var x = 3
	fmt.Println(x)
	fmt.Println(factorial(8))
	array_demo()
	fmt.Println("substring of 'abcde' from 1 to 3 is:", substring(1, 3, "abcde"))
	fmt.Println("substring2 of 'abcde' from 1 to 3 is:", substring2(1, 3, "abcde"))
	global := 43
	switch_case(mytype(24))

	switch_case(24)
	switch_case(48.0)
	fmt.Println("local global is ", global)
	a := "abc"
	fmt.Println(a == "abc")
	double_dimensional_array(3, 5)
	inner_func()
	set_demo()
	map_demo()
	inner_recur_demo()
	oop_demo()
	string_builder_demo()
	string_convert_demo()
}
