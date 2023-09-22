package main

import "fmt"

// Ushbu misolda taxmin qilingan uzunliklarga ega ikkita massiv (arr1 va arr2) e'lon qilinadi:
func main() {
	var arr1 = [...]int{5, 6, 7, 8, 9} // yani bunda [...] berish orqali arraysga n ta malumot yuklashimiz mumkin

	fmt.Println(len(arr1))

	// go Slices orqali array malumotlarni olish
	myslice := arr1[0:2] // 0 dan 2 gacha malumotlarni olish uchun Slice ishlatilishi

	fmt.Println(myslice)

}
