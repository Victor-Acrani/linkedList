package main

import (
	"fmt"
	"strings"
)

// defining node struct
type node struct {
	value int
	next  *node
}

// defining linkList struct
type linkedList struct {
	head   *node
	length int
}

func main() {
	//creating linkedList
	linkedList := linkedList{}
	//adding values
	linkedList.add(1)
	linkedList.add(2)
	linkedList.add(3)
	linkedList.add(4)
	linkedList.add(5)

	//printing whole list
	fmt.Println(linkedList)
	//getting value
	fmt.Println(linkedList.get(3))
	//removing values
	linkedList.remove(3)
	linkedList.remove(5)
	//printing list after removing values
	fmt.Println(linkedList)
	//removing first node
	linkedList.remove(1)
	//printing list after removing values
	fmt.Println(linkedList)
}

func (l *linkedList) add(value int) {
	//cerating new node and defining value
	newNode := new(node)
	newNode.value = value

	//checking if linkedList is empty or not
	if l.head == nil {
		l.head = newNode //pass the node reference to the head
	} else {
		iterator := l.head
		for ; iterator.next != nil; iterator = iterator.next { //get the last node
		}
		iterator.next = newNode //pass the node reference
	}
}

func (l *linkedList) remove(value int) {
	// pointer to save the previous node
	var previous *node

	for current := l.head; current != nil; current = current.next {
		if current.value == value {
			if current == l.head { //case the node to remove is the first of the linkedList
				l.head = current.next
			} else { //other nodes
				previous.next = current.next
				return
			}
		}
		previous = current
	}
}

func (l linkedList) get(value int) *node {
	//check if value exists
	for iterator := l.head; iterator != nil; iterator = iterator.next {
		if iterator.value == value {
			return iterator
		}
	}
	return nil
}

func (l linkedList) String() string {
	stringBuilder := strings.Builder{}

	//formatting how the linkedList will be displayed
	for iterator := l.head; iterator != nil; iterator = iterator.next {
		stringBuilder.WriteString(fmt.Sprintf("%s", iterator))
	}
	return stringBuilder.String()
}

func (n node) String() string {
	//formatting how the node will be displayed
	return fmt.Sprintf("%d", n.value)
}
