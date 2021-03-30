package list

import "strconv"

// Node represents a linked list node
type Node struct {
	val  int
	next *Node
}

// List represents a linked list data stucture
type List struct {
	head *Node
	tail *Node
	len  int
}

// NewNode creates a new node
func NewNode(val int, next *Node) *Node {
	return &Node{val: val, next: next}
}

//New creates a new list
func New() *List {
	tail := NewNode(0, nil)
	head := NewNode(0, tail)
	return &List{head: head, tail: tail}
}

// AddToFront adds a node to the front of the list
func (l *List) AddToFront(val int) {
	node := NewNode(val, l.head.next)
	l.head.next = node
	l.len++
}

// RemoveFromFront removes a node from the front of the list
func (l *List) RemoveFromFront() {
	if l.len != 0 {
		l.head.next = l.head.next.next
		l.len--
	}
}

// AddToBack adds a node to the back of the list
func (l *List) AddToBack(val int) {
	newTail := NewNode(0, nil)
	l.tail.val = val
	l.tail.next = newTail
	l.tail = newTail
	l.len++
}

// String method to print contents of List
func (l *List) String() string {
	var listStr string
	head := l.head.next
	for head.next != nil {
		if listStr != "" {
			listStr += "->"
		}
		listStr += strconv.Itoa(head.val)
		head = head.next
	}
	return listStr
}
