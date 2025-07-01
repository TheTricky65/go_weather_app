package main

import (
	"fmt"
)


func main() {


myArray := [3]int{1, 2, 3}
for i, v := range myArray {
	fmt.Printf("Index: %d, Value: %d\n", i, v)
}
var test = 10
for i:= range test{

fmt.Printf("number is %d\n", i)

}

}