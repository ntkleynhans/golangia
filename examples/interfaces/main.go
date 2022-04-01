package main

import "fmt"
import "interfaces/shape"

//type Circle struct {
//	name string
//}
var print = fmt.Println

func main() {
	circle := shape.Circle{"circle"}
	var item shape.Shape
	item = &circle
	print(item.Name())
}
	
