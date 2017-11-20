package lists

import (
	"fmt"
)

type node struct {
	data interface{}
	next *node
}

type List struct {
	head *node
	len  int
}

func Helloworld() bool {
	fmt.Println("hello world")
	return true
}

func NewList() (*List, bool) {
	l := new(List)

	if l == nil {
		fmt.Println("New list create failed")
		return nil, false
	}
	return l, true
}

func (l *List) Addtolist(d interface{}) {
	l.head = &node{d, l.head}
	l.len++
}

func (l *List) Removefromlist(d interface{}) bool {
	if l.len > 0 {
		n := l.head
		var prev *node = nil

		for n != nil {
			if n.data == d {
				if prev != nil {
					prev.next = n.next
					l.len--
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

func (l *List) Print() (int, error) {

	n := l.head

	for n != nil {
		fmt.Println(n.data)
		n = n.next
	}

	return l.len, nil
}
