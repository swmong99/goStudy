package main

import "fmt"

func main(){
	var num1 *int = new(int)
	var num2 *int = new(int)

	fmt.Scanln(num1, num2)
	fmt.Println(*num1, *num2)
	swap(num1, num2)
	fmt.Println(*num1, *num2)

}

func swap(a *int, b *int){
	var c int
	c = *a
	*a = *b
	*b = c
}
