package main

import "fmt"

func main() {
	var a, b, c int = 2, 3, 4
	var name, my_var, myvar string = "golang", "jhon", "euro"

	fmt.Println(name,
		my_var,
		myvar,
	)
	fmt.Println(a, b, c)

	// bush qiymat bulsa default 0 ni olib ketadi masalan --> var a int
	var (
		an int
		bn int    = 1
		cn string = "hello"
	)

	fmt.Println(an)
	fmt.Println(bn)
	fmt.Println(cn)
}

// main func println
//func main() {
//	// Variables yozilish usulli 1
//	var name = "matn" // string
//	var son = 34      // int
//	var soon = 3.0    // float32
//	var bin = true    // boolean
//
//	// Variables yozilish usulli 2
//	name1 := "matn" // string
//	son1 := 34      // int
//	soon1 := 3.0    // float32
//	bin1 := true    // boolean
//
//	// Variables yozilish usulli 3
//	var my_var string
//	my_var = "test text"
//	var myvar int
//	myvar = 50
//	fmt.Println(myvar)
//	fmt.Println(my_var)
//
//	// chop etish uchun
//	fmt.Println("hello world")
//	fmt.Println(name)
//	fmt.Println(son)
//	fmt.Println(soon)
//	fmt.Println(bin)
//	fmt.Println(name1)
//	fmt.Println(son1)
//	fmt.Println(soon1)
//	fmt.Println(bin1)
//
//}
