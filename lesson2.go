package main

import "fmt"

// constanta o'zgarmas o'zgaruvchilar
func main() {

	// still 1
	const pi = 3.14
	// still 2
	const (
		a int = 14
		b int = 34
	)
	fmt.Println(pi, a, b)
}
