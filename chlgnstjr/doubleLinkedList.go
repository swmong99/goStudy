// double project main.go
package main

import (
   "fmt"
)

type dlinkedlist struct {
   next *dlinkedlist
   prev *dlinkedlist
   data int
}

func main() {
   var headList *dlinkedlist = new(dlinkedlist)
   var tailList *dlinkedlist = new(dlinkedlist)
   var printList *dlinkedlist = new(dlinkedlist)

   add_append(headList, tailList, 1)
   add_append(headList, tailList, 2)
   add_append(headList, tailList, 3)
   add_append(headList, tailList, 4)
   add_append(headList, tailList, 5)
   add_append(headList, tailList, 6)

   printList.next = headList.next

   for {
      fmt.Println(printList.next.data)
      printList.next = printList.next.next
      if printList.next == nil {
         break
      }
   }
   printList.next = tailList.next
   for {
      fmt.Println(printList.next.data)
      printList.next = printList.next.prev
      if printList.next == nil {
         break
      }
   }

}

func add_append(headList *dlinkedlist, tailList *dlinkedlist, t int) {
   var appendList *dlinkedlist = new(dlinkedlist)

   appendList.data = t

   if headList.next == nil {
      headList.next = appendList
      tailList.next = appendList
   } else {
      tailList.next.next = appendList
      appendList.prev = tailList.next /* 추가*/
      tailList.next = appendList
   }

}
