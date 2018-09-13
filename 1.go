package main

import "fmt"
import _ "os"
import _ "runtime" 

func ready_input()(int){
   fmt.Println("input number")
   var input int
   fmt.Scanln(&input)
   return input
}


func cnt_element(x int){
 if x < 1 || x>= 10 {
   fmt.Println("plz input only 2 to 9 !!!")
   panic("error!!")
  } else {

   twoD := make([][]int ,10)

   for i:=1 ; i< 10 ; i++ {
    if i==1 {
     twoD[i] = make([]int , 1)
     twoD[i][0] = x
     arr(twoD[i])
    } else {
    fmt.Println(x)
    }
   } 

  }
 fmt.Print(x)
}

func arr(a []int){
  fmt.Println(a)
}


func f() {
   defer func() {
      s := recover()
      fmt.Println(s)
   }()
}

func main() {
 f()
 num:=ready_input() 
 cnt_element(num)
 fmt.Println("end")
 
}
