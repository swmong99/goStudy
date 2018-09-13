package main

import "fmt"

func main () {
    recur:
    var n1 int 
    fmt.Print("숫자를 입력하세요 (종료는 0) : ")
    fmt.Scan(&n1)
        if n1 == 0 {
            fmt.Println("프로그램을 종료합니다")
        } else if n1 > 1 && n1 < 10 {
            for i := 1; i < 10; i++ {
                fmt.Println(n1,"*",i,"=",n1*i)
            }
            goto recur
        } else {
            fmt.Println("2 ~ 9 사이의 값을 입력해 주세요.")
            goto recur    
    }
}

