package main

import "fmt"

func main() {
	var s1 string = "한글"
	var s2 string = "한글"
	var s3 string = "Go"

	fmt.Println(s1 == s2)     // true: 두 문자열이 같으므로 true
	fmt.Println(s1 != s2)     // false: 두 문자열이 다르므로 false
	fmt.Println(s1 + s3 + s2) // 한글Go한글
	fmt.Println("안녕하세요" + s1) // 안녕하세요한글

	var s4 string = "1234567890"

	for i:=0; i<len(s4); i++ {
		fmt.Println(s4[i])
		fmt.Printf("%c\n", s4[i])
	}


}
