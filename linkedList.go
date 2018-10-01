package main

import "fmt"

type Node struct{
	body int
	next *Node
}

func Append(body int, head *Node, tail *Node){
	var newNode *Node = new(Node)

	newNode.body = body
	newNode.next = nil

	if head.next == nil {
		head.next = newNode
		tail.next = newNode
	} else {
		tail.next.next = newNode
		tail.next = newNode
	}
}

func main(){

	var headNode *Node = new(Node)
	var tailNode *Node = new(Node)

	Append(1, headNode, tailNode)
	Append(2, headNode, tailNode)
	Append(3, headNode, tailNode)
	Append(4, headNode, tailNode)

	var currNode *Node = headNode.next

	for {
		fmt.Println(currNode.body)

		if currNode.next == nil{
			break
		 } else{
		 	currNode = currNode.next
		}

	}
}

