package main

import (
	"fmt"
	"reflect"
)

func main(){
	for{
		fmt.Print("숫자를 입력하세요(종료는 0) : ")
		var num int

		// Scanln의 리턴값을 이용하여 타입 체크를 한다
		_, err := fmt.Scanln(&num)

		// err 변수가 nil일때 정상
		if err != nil {
			fmt.Println("입력 오류입니다. 정수형 변수를 입력하세요.")

			// 문자'열' 입력시 버퍼를 비우기 위함
			var clearBuffer string
			fmt.Scanln(&clearBuffer)
		} else if num == 0 {
			fmt.Println("프로그램을 종료합니다.")
			break
		} else if num < 1 || num > 9 {
			fmt.Println("1 ~ 9사이의 값을 입력해주세요.")
		} else {	//여기서부터 정상처리
			// 2차원 슬라이스를 생성
			var sliceCnt [][]int
			// 초기 생성크기는 1개로 설정한다
			sliceCnt = make([][]int, 1)	// 1차원 배열 메모리 할당
			sliceCnt[0] = make([]int, 1) // 2차원 배열 메모리 할당

			// 입력값을 초기 2차원 배열 [0][0]에 입력
			sliceCnt[0][0] = num

			fmt.Println(sliceCnt[0])

			var i int = 0
			//tempSliceCnt := make([]int, 2)
			for {
				var arrSize int
				if i == 0 {
					arrSize = 2
				} else{
					arrSize = len(sliceCnt[i])
				}
				tempSliceCnt := make([]int, arrSize)
				for j := range sliceCnt[i]{
					if len(tempSliceCnt) <= ((sliceCnt[i][j] - 1) * 2){
						tempSliceCnt = append(tempSliceCnt, sliceCnt[i][j])
						tempSliceCnt = append(tempSliceCnt, 1)
					} else{
						tempSliceCnt[(sliceCnt[i][j] - 1) * 2] = sliceCnt[i][j]
						tempSliceCnt[((sliceCnt[i][j] - 1) * 2) + 1] += 1
					}

				}
				fmt.Println(tempSliceCnt)
				sliceCnt = append(sliceCnt, tempSliceCnt)
				i += 1

				if i >= 2{
					if reflect.DeepEqual(sliceCnt[i], sliceCnt[i-1]){
						if reflect.DeepEqual(sliceCnt[i-1], sliceCnt[i-2]) {
							fmt.Println("같은값이 반복되어 종료합니다.")
							break
						}
					}
				}
			}
		}
	}
}