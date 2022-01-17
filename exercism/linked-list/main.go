package main

import "fmt"

type node struct {
	next *node
	val  int
}

type LinkedList struct {
	head   *node
	length int
}

func (l *LinkedList) prepend(n *node) {

	/*In order to add something to the beginning of the linked list,
	we need to first create a temporary variable called second to store
	the value of the head as we want to add something in the
	first node or head. */
	second := l.head

	/*This sets the new node, n, as the head which means it'll be at the
	beginning of the linked list. */
	l.head = n

	/*We then link our list back together by assigning the address of the
	second to the node n which is now in first place or the head*/
	l.head.next = second

	/*We then need to increase the length of the linked list as we added
	something to it*/
	l.length++
}

func (l LinkedList) printListData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.val)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Println("")

}

func main() {

	/*The line below creates a new linked list*/
	mylist := LinkedList{}

	/*The line below creates a node, to create a node you need to pass
	in a pointer hence the '&'. */
	node1 := &node{val: 100}
	node2 := &node{val: 200}
	node3 := &node{val: 300}
	node4 := &node{val: 400}

	mylist.prepend(node1)
	mylist.prepend(node2)
	mylist.prepend(node3)
	mylist.prepend(node4)

	/*The out come of this print call = {0xc000010090 4}.
	The output is the memory location of the 4th node which we
	added last in this linked and is thus the head of the list
	with 4 being the length of mylist.*/
	fmt.Println(mylist)

	mylist.printListData()

}
