package main

import "fmt"

type node struct {
	data int16
	next *node
}

type linkedList struct {
	head   *node
	length int16
}

func (l *linkedList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second

	l.length++
}

func (l linkedList) printListDate() {
	toPrint := l.head // printing head first
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}

	fmt.Printf("\n")
}

func (l *linkedList) deleteWithValue(value int16) {
	// handling empty list
	if l.length == 0 {
		return
	}
	// handling delete header. second node should become the first node
	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}
	previousToDelete := l.head

	for previousToDelete.next.data != value {
		if previousToDelete.next.next == nil { // handling if a number is not found in linkedlist .so, we will not have any nil pointer dereference
			return
		}
		previousToDelete = previousToDelete.next
	}

	previousToDelete.next = previousToDelete.next.next

	l.length--
}

func main() {
	myList := &linkedList{}
	node1 := &node{data: 48}
	node2 := &node{data: 8}
	node3 := &node{data: 4}
	node4 := &node{data: 548}
	node5 := &node{data: 778}
	node6 := &node{data: 884}
	node7 := &node{data: 3848}
	node8 := &node{data: 3838}
	node9 := &node{data: 3874}
	node10 := &node{data: 4128}
	node11 := &node{data: 8137}
	node12 := &node{data: 4303}

	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)
	myList.prepend(node4)
	myList.prepend(node5)
	myList.prepend(node6)
	myList.prepend(node7)
	myList.prepend(node8)
	myList.prepend(node9)
	myList.prepend(node10)
	myList.prepend(node11)
	myList.prepend(node12)

	// fmt.Println(*myList)
	myList.printListDate()
	myList.deleteWithValue(100) //
	// To handle run time error
	//  handle delete header
	// and unknown number needs to be handled.
	// delete from empty list

	myList.printListDate()

	emptyList := linkedList{}

	emptyList.deleteWithValue(100)
}
