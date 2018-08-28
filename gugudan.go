package main

import "fmt"

func main() {

	for{
		fmt.Print("숫자를 입력하세요(종료는 0) : ")
		var num int

		// Scanln의 리턴값을 이용하여 타입 체크를 한다
		_, err := fmt.Scanln(&num)

		if err != nil {
			fmt.Println("입력 오류입니다. 정수형 변수를 입력하세요.")

			// 문자'열' 입력시 버퍼를 비우기 위함
			var clearBuffer string
			fmt.Scanln(&clearBuffer)
		} else if num == 0 {
			fmt.Println("프로그램을 종료합니다.")
			break
		} else if num < 2 || num > 9 {
			fmt.Println("2 ~ 9사이의 값을 입력해주세요.")
		} else {
			for i := 1; i<=9; i++ {
				fmt.Printf("%d * %d = %d\n", num, i, num * i)
			}
		}

	}
}

