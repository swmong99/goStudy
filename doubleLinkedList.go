package main

import "fmt"

type Node struct {
	prev *Node
	body int
	next *Node
}

func Append(body int, head *Node, tail *Node) {
	var newNode *Node = new(Node)

	newNode.body = body
	newNode.next = nil
	newNode.prev = nil

	if head.next == nil {
		head.next = newNode
		head.prev = nil
		
		tail.next = newNode
		tail.prev = nil
	} else {
		newNode.prev = tail.next
		tail.next.next = newNode
		tail.prev = tail.next
		tail.next = newNode
		
	}
}

func main() {

	var headNode *Node = new(Node)
	var tailNode *Node = new(Node)

	Append(1, headNode, tailNode)
	Append(2, headNode, tailNode)
	Append(3, headNode, tailNode)
	Append(4, headNode, tailNode)

	/*
	var currNode *Node = headNode.next

	for {
		fmt.Println(currNode.body)

		if currNode.next == nil {
			break
		} else {
			currNode = currNode.next
		}

	}
	*/
	
	var currNode *Node = tailNode.next

	for {
		fmt.Println(currNode.body)

		if currNode.prev == nil {
			break
		} else {
			currNode = currNode.prev
		}

	}
}

