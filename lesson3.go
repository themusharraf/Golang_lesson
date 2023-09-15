package main

import "fmt"

// go lang arrays
//func main() {
//	arr := [3]int{1, 2, 3}
//	arr2 := [5]string{"hello", "test", "go lang"}
//	arr3 := [4]string{"Tesla", "BMW", "Mazda", "Ford"}
//	fmt.Println(arr[0]) // yani arr[0] index da turgan malumotni chiqarish uchun
//	fmt.Println(arr2[1])
//	fmt.Println(arr3[2])
//}

//func main() {
//	arr4 := [2]int{} // bu holatda default 0 bilan array ni tuldiradi
//	arr5 := [3]int{1, 2, 3}
//	arr6 := [4]int{1, 2, 3, 4}
//
//	fmt.Println(arr4)
//	fmt.Println(arr5)
//	fmt.Println(arr6)
//}

func main() {
	arr := [3]int{20, 30, 40}
	arr[1] = 50 // yani bu yerda array ni 1 index ni o'zgartirdim
	fmt.Println(arr)
}
