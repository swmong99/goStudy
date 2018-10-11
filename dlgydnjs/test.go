// struct
package main

import (
   "fmt"
)

type NODE struct {
   prev *NODE
   next *NODE
   data int
}

func addnode(a *NODE, b *NODE, c int) {
   new_node := new(NODE)
   new_node.data = c

   if a.prev == nil {
      a.prev = new_node
      a.next = new_node
      b.prev = new_node
      b.next = new_node
   } else {
      a.next = b.prev
      b.prev = b.next
      b.next = new_node
   }
}

func main() {
   before := new(NODE)
   after := new(NODE)

   addnode(before, after, 10)
   addnode(before, after, 20)
   addnode(before, after, 30)
   addnode(before, after, 40)
   addnode(before, after, 50)
   addnode(before, after, 60)
   addnode(before, after, 70)
   addnode(before, after, 80)
   addnode(before, after, 90)


   var bf_pr *NODE = before.prev
   var bf_nx *NODE = before.next
   var af_pr *NODE = after.prev
   var af_nx *NODE = after.next

   fmt.Println(bf_pr.data)
   fmt.Println(bf_nx.data)
   fmt.Println(af_pr.data)
   fmt.Println(af_nx.data)
}
