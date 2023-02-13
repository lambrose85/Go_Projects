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

func (L *Linked) addNode(val int) {
	newNode := Node{data: val}
	if L.head != nil { // checking to see if there is currently a head node
		newNode.next = L.head
		L.head = &newNode
		L.length++
	} else { // creates head node in case the list is empty
		L.head = &newNode
		L.length++
	}
}

func (L *Linked) printList() {
	if L.head == nil { //checking for nodes in list
		fmt.Println("empty list")
	} else {
		currentNode := L.head // sets current node equal to the head of the list
		for currentNode != nil {
			fmt.Println(currentNode.data)
			currentNode = currentNode.next
		}
	}

}
