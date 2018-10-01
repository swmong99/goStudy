package main

import "fmt"

type Node struct{
	body int
	next *Node
}

func Append(body int, head *Node, curr *Node){
	var newNode *Node = new(Node)

	if head == nil {
		head = new(Node)
		head.next = newNode
	}

	newNode.body = body
	newNode.next = nil

	curr = newNode
}

func main(){

	var headNode *Node
	var currNode *Node = new(Node)
	Append(1, headNode, currNode)

	fmt.Println(currNode.body)
}
