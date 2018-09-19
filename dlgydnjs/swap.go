package main

 

import "fmt"

 

 

func swap(a *int, b *int){

    var c int

 

    c = *a

   *a = *b

   *b = c

 

}

 

 

func main() {

    first:

    var u1 *int = new(int)

    fmt.Print("첫번째 값을 입력하세요:")

    _, err := fmt.Scanln(u1)

    if err != nil {

                    fmt.Println("첫번째 값에 숫자를 입력 하세요 !!!")

                    var cla1 string

                    fmt.Scanln(&cla1)

                    goto first

    }

 

    second:

    var u2 *int = new(int)

    fmt.Print("두번째 값을 입력하세요:")

    _, err1 := fmt.Scanln(u2)

    if err1 != nil {

                    fmt.Println("두번째 값에 숫자를 입력하세요 !!!")

                    var cla2 string

                    fmt.Scanln(&cla2)

                    goto second

    }

   

    

    swap(u1,u2)

    fmt.Println("첫번째 값은:",*u1,"두번째 값은:",*u2,"# 체인지 성공!!!!!!!")

}
