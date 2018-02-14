package lists

import (
	"fmt"
)

type node struct {
	data interface{}
	next *node
}

//Linked List Struct
type List struct {
	head *node
	size int
}

//Create list
func New() (*List, bool) {
	l := new(List)

	if l == nil {
		fmt.Println("New list create failed")
		return nil, false
	}
	return l, true
}

func (n *node) searchnode(d interface{}) bool {
	if n == nil {
		return false
	} else if n.data == d {
		return true
	} else {
		return n.next.searchnode(d)
	}
}

//Search the list
func (l *List) Search(d interface{}) bool {
	return l.head.searchnode(d)
}

//Insert to the List
func (l *List) Insert(d interface{}) bool {

	l.head = &node{d, l.head}
	l.size++
	return true
}

//Delete from the list
func (l *List) Delete(d interface{}) bool {
	if l.size > 0 {
		n := l.head
		var prev *node

		for n != nil {
			if n.data == d {
				if prev != nil {
					prev.next = n.next
					l.size--
				} else {
					l.head = n.next
				}
				n = nil
				return true
			}
			prev, n = n, n.next
		}
	}
	return false
}

//Print the list
func (l *List) Print() (int, error) {

	if l == nil {
		return 0, nil
	}

	n := l.head

	for n != nil {
		fmt.Println(n.data)
		n = n.next
	}

	return l.size, nil
}
