package linklist

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) Append(data int) {
	newNode := &Node{data: data}
	if list.head == nil {
		list.head = newNode
		return
	}
	lastNode := list.head
	for lastNode.next != nil {
		lastNode = lastNode.next
	}
	lastNode.next = newNode
}

func (list *LinkedList) PrintList() {
	currentNode := list.head
	for currentNode != nil {
		fmt.Printf("%d ->", currentNode.data)
		currentNode = currentNode.next
	}
	fmt.Println("nil")
}
