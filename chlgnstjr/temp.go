package main

import "fmt"
import "os"
import _ "runtime" 

func ready_input(){
   fmt.Println("input number")
   var input int
   fmt.Scanln(&input)
   googd(input)
}

func googd(x int) {
  if x == 1 || x>= 10 {
   fmt.Println("plz input only 2 to 9 !!!")
   ready_input()  
  } else if x == 0  {
   fmt.Println("program exit")
   os.Exit(1)
  } else {
   googd_print(x)
  }
}

func googd_print(x int){
 for j:=1; j<10; j++ {
  fmt.Printf("%d * %d = %d\n" ,x,j,x*j )
 }
}

func main() {
ready_input()  
}
