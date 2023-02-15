package main

/*

  head                                                                             tail

 | 0 | -> | 1 | -> | 2 | -> | 3 | -> | 4 | -> | 5 | -> | 6 | -> | 7 | -> | 8 | -> | 9 |

 | 9 | -> | 8 | -> | 7 | -> | 6 | -> | 5 | -> | 4 | -> | 3 | -> | 2 | -> | 1 | -> | 0 |

*/

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

// Simple reverse a singly linked list
func reverse_list(head *Node) *Node {
	cur := head
	var prev *Node = nil
	var next *Node = nil
	for cur != nil {
		next = cur.next
		cur.next = prev
		prev = cur
		cur = next
	}
	return prev
}

func print_list(node *Node) {
	for node != nil {
		fmt.Printf("%v\n", *node)
		node = node.next
	}
}

func main() {
	head := &Node{0, nil} // Looks like "new" in C++, but it does not require "delete". The garbage collector takes care of it.
	next := head
	for i := 1; i < 10; i++ {
		next.next = &Node{i, nil}
		next = next.next
	}

	print_list(head)
	head = reverse_list(head)
	print_list(head)

}
