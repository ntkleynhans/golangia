package main

import "fmt"

var print = fmt.Println

func Add(a int, b int) (res int) {
	return a + b
}

func main() {
	print(Add(1,4))
}
	
