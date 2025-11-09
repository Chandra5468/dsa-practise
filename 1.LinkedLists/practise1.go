package main

import "fmt"

type node_test struct {
	next_node_test *node_test
	data           int16
}

type linked_list_test struct {
	head_node *node_test
	length    int16
}

func (l *linked_list_test) prepend_test(n *node_test) { // prepending a node to start
	second := l.head_node
	l.head_node = n
	l.head_node.next_node_test = second

	l.length++

}

func testmain() {
	// creating 3 or 4 nodes

	myList := &linked_list_test{}

	node1 := &node_test{data: 48}
	node2 := &node_test{data: 8}
	node3 := &node_test{data: 4}

	myList.prepend_test(node1)
	myList.prepend_test(node2)
	myList.prepend_test(node3)

	fmt.Println(myList)
}
