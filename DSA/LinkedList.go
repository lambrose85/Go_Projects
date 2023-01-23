package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type Linked struct {
	head   *Node
	length int
}

func main() {
	newList := Linked{}
	newList.addNode(10)
	newList.addNode(25)
	newList.addNode(1000)
	newList.printList()
}

func (L *Linked) addNode(x int) {
	newNode := Node{data: x}
	if L.head != nil {
		newNode.next = L.head
		L.head = &newNode
		L.length++
		fmt.Println("added to list")
	} else {
		L.head = &newNode
		L.length++
		fmt.Println("added to list")
	}

}
func (L *Linked) printList() {
	if L.head == nil {
		fmt.Println("Empty list")
		return
	}
	currentNode := L.head
	for currentNode != nil {
		fmt.Println(currentNode.data)
		currentNode = currentNode.next
	}
}
