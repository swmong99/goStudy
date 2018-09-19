package main

import "fmt"
import _"runtime" 

func main() {
 var fstptr *string = new(string)
 var sndptr *string = new(string)

 fmt.Println("Plz input first parameter")
 fmt.Scanln(fstptr)

 fmt.Println("Plz input second parameter")
 fmt.Scanln(sndptr)

 swap(fstptr ,sndptr )

 fmt.Println(*fstptr) 
 fmt.Println(*sndptr)
}

func swap(aa *string ,bb *string ){
 var temp *string = new(string)
 *temp = *aa
 *aa =*bb 
 *bb = *temp
}
 
