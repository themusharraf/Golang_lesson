package main

import "fmt"

// Ushbu misolda taxmin qilingan uzunliklarga ega ikkita massiv (arr1 va arr2) e'lon qilinadi:
func main() {
	var arr1 = [...]int{5, 6, 7, 8, 9} // yani bunda [...] berish orqali arraysga n ta malumot yuklashimiz mumkin
	arr := [...]int{1, 3, 4, 5, 6, 7}
	fmt.Println(len(arr))
	fmt.Println(len(arr1))

	// go Slices orqali array malumotlarni olish
	myslice := arr1[0:2]

	fmt.Println(myslice)

}
