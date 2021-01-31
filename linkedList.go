package main

import "fmt"

type Node struct {
	Data string
	Next *Node
}

type LinkedList struct {
	Length int
	Head   *Node
}

func main() {
	var list LinkedList

	node1 := Node{Data: "Node one"}
	list.insertNode(&node1)

	node2 := Node{Data: "Node two"}
	list.insertNode(&node2)

	node3 := Node{Data: "Node three"}
	list.insertNode(&node3)

	list.printData()

	list.deleteNode("sdgsdgdsg")
	fmt.Println()

	list.deleteNode("Node two")
	list.printData()

	list.deleteNode("Node one")
	list.printData()

}

// Inserts a node into a LinkedList
func (list *LinkedList) insertNode(newNode *Node) {
	if list.Length == 0 {
		list.Head = newNode
	} else {
		currentNode := list.Head
		for currentNode.Next != nil {
			currentNode = currentNode.Next
		}
		currentNode.Next = newNode
	}

	list.Length++
}

// Removes a node into a LinkedList
func (list *LinkedList) deleteNode(query string) {
	var previousNode *Node

	if list.Head.Data == query {
		list.Head = list.Head.Next
		return
	}

	currentNode := list.Head
	for currentNode.Next != nil {
		// if data matches query point the previous node to the next node
		if currentNode.Data == query {
			if previousNode != nil {
				previousNode.Next = currentNode.Next
			}

			list.Length--
			fmt.Println("Deleted!")
			fmt.Println("List Length: ", list.Length)
			fmt.Println()
			return
		}
		previousNode = currentNode
		currentNode = currentNode.Next
	}

	fmt.Println("Could not find: " + query)

}

// Print data from all nodes in a LinkedList
func (list *LinkedList) printData() {
	currentNode := list.Head

	for currentNode.Next != nil {
		fmt.Println(currentNode.Data)
		currentNode = currentNode.Next
	}

	fmt.Println(currentNode.Data)
	fmt.Println()
}
