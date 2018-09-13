package main

import "fmt"

func main () {
    var n1   int
    var cnt1 int
    var cnt2 int
    var cnt3 int
    var cnt4 int
    var cnt5 int
    var cnt6 int
    var cnt7 int
    var cnt8 int
    var cnt9 int
    var s1 []int
    var s2 []int = make([]int, 2)
    var s3 []int
    var s4 []int
    var s5 []int
    var s6 []int
    var s7 []int
    var s8 []int
    var s9 []int

    fmt.Print("1을 입력하세요 : ")
    fmt.Scan(&n1)
    s1 = append(s1,n1)
 
    for i:=0; i < len(s1); i++ {
    chk := s1[i]
    switch chk {
    case 1:
        cnt1 += 1
    case 2:
        cnt2 += 1
    case 3:
        cnt3 += 1
    case 4:
        cnt4 += 1
    case 5:
        cnt5 += 1
    case 6:
        cnt6 += 1
    case 7:
        cnt7 += 1
    case 8:
        cnt8 += 1
    case 9:
        cnt9 += 1
        }
    }
    s1 = append(s1,cnt1)
    s2 = append(s2,cnt2)
    s3 = append(s3,cnt3)
    s4 = append(s4,cnt4)
    s5 = append(s5,cnt5)
    s6 = append(s6,cnt6)
    s7 = append(s7,cnt7)
    s8 = append(s8,cnt8)
    s9 = append(s9,cnt9)
    fmt.Println(s1)
    copy(s2,s1)

    fmt.Println(s2)
    for i:=0; i < len(s2); i++ {
    chk2 := s2[i]
    switch chk2 {
    case 1:
        cnt1 += 1
    case 2:
        cnt2 += 1
    case 3:
        cnt3 += 1
    case 4:
        cnt4 += 1
    case 5:
        cnt5 += 1
    case 6:
        cnt6 += 1
    case 7:
        cnt7 += 1
    case 8:
        cnt8 += 1
    case 9:
        cnt9 += 1
            }
    }
    s1 = append(s1,cnt1)
    s2 = append(s2,cnt2)
    s3 = append(s3,cnt3)
    s4 = append(s4,cnt4)
    s5 = append(s5,cnt5)
    s6 = append(s6,cnt6)
    s7 = append(s7,cnt7)
    s8 = append(s8,cnt8)
    s9 = append(s9,cnt9)
    fmt.Println(s2)
    copy(s3,s2)
}
