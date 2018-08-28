package main

import (
	"fmt"
	"reflect"
)

func main() {

	for{
		fmt.Print("숫자를 입력하세요(종료는 0) : ")
		var num int

		fmt.Scanln(&num)
		//fmt.Println(num)
		//fmt.Println(reflect.TypeOf(num).String())

		if num == 0{
			fmt.Println("프로그램을 종료합니다.")
			break
		} else if num < 2 || num > 9 {
			fmt.Println("2 ~ 9사이의 값을 입력해주세요.")
		} else if reflect.TypeOf(num).String() != "int" {
			fmt.Println("2 ~ 9사이의 값을 입력해주세요.")
		} else {
			for i := 1; i<=9; i++ {
				fmt.Printf("%d * %d = %d\n", num, i, num * i)
			}
		}

	}
}

