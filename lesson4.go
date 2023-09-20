package main

import "fmt"

func main() {
	var arr1 = [...]int{5, 6, 7, 8, 9}
	arr := [...]int{1, 3, 4, 5, 6, 7}
	fmt.Println(len(arr))
	fmt.Println(len(arr1))
}
