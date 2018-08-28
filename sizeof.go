package main

import "fmt"
import "unsafe" // Sizeof 함수를 사용하기 unsafe 패키지 사용

func main() {
	var num1 int8  = 1
	var num2 int16 = 1
	var num3 int32 = 1
	var num4 int64 = 1
	var num5 int = 1
	var name1 string = "leehochul"
	var name2 string = "abc"
	var name3 string = "1234567890123456789012345678901234567890"
	var name4 = "1234567890"

	fmt.Println(unsafe.Sizeof(num1)) // 1
	fmt.Println(unsafe.Sizeof(num2)) // 2
	fmt.Println(unsafe.Sizeof(num3)) // 4
	fmt.Println(unsafe.Sizeof(num4)) // 8
	fmt.Println(unsafe.Sizeof(num5)) // 8, 64bit system
	fmt.Println(unsafe.Sizeof(name1))
	fmt.Println(unsafe.Sizeof(name2))
	fmt.Println(unsafe.Sizeof(name3))
	fmt.Println(len(name3))
	fmt.Println(len(name4))
	fmt.Println(unsafe.Sizeof(name4))

}
