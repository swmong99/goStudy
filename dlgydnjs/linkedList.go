package main

import "fmt"


type NODE struct {
    no_add *int
    data int
}

func addnode(c *NODE int,b int) {
    var a int = *c 
    new_node := new(NODE)
    new_node.no_add = &a
    new_node.data = b

    c.no_add = new_node   
}

func main() {
    new_node1 := new(NODE)
    new_node1.no_add = new(int)
    new_node1.no_add = nil
    
    chk := addnode(&new_node1,10)
    if chk != nil {
            fmt.Scanln(new_node1)
    }
}
