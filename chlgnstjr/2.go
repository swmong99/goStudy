package main

import "fmt"
import _ "os"
import _ "runtime" 



func cnt_element(x int){
 var s []int 
 var cnt1,cnt2,cnt3,cnt4,cnt5 int =0,0,0,0,0
 s = append(s,x  )

 if x < 1 || x> 1 {
   fmt.Println("plz input only 1 !!!")
   panic("error!!")
  } else {

   loop :
   for i :=0 ; i< len(s) ; i=i+1 { 
    switch s[i]  {
    case 1:
      cnt1++
    case 2:
      cnt2++
    case 3:
      cnt3++
    case 4:
      cnt4++
    case 5:
      cnt5++
    default :
    }
    if i == len(s)-1 {
       var ss []int
     

        if cnt1 != 0 {
         ss = append(ss, cnt1 ) 
        } else if cnt2 != 0 {
         ss = append(ss, cnt2 ) 
        } else if cnt3 != 0 {
         ss = append(ss, cnt3 ) 
        } else if cnt4 != 0 {
         ss = append(ss, cnt4 ) 
        } else if cnt5 != 0 { 
         ss = append(ss, cnt5 ) 
        }

     fmt.Println(ss)

     s=ss    
     goto loop

    } else {
     continue
    }  


   
  }
  fmt.Print(x)
}


func f() {
   defer func() {
      s := recover()
      fmt.Println(s)
   }()
}

func main() {
 f()
   fmt.Println("input number")
   var input int
   fmt.Scanln(&input)

 cnt_element(input)
 fmt.Println("end")
 
}
