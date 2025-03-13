package main

import (
	"fmt"
)

func main() {

	// fmt.Println(runtime.NumCPU())

	defer fmt.Println("Это выполняется последним")
	fmt.Println("Это выполняется первым")

	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("Main завершён")
}
