// dbl-link-go, LitFill <marrazzy54 at gmail dot com>
// program for training to implement doubly linked list in go.
package main

import (
	"fmt"
	"strings"
)

type Node struct {
	next, before *Node
	data         string
}

func NewNode(data string) *Node {
	return &Node{data: data}
}

func NewNodes(data []string) (head *Node) {
	if len(data) == 0 {
		return nil
	}
	head = NewNode(data[0])
	current := head
	for _, d := range data[1:] {
		node := NewNode(d)
		current.setNext(node)
		current = node
	}
	return
}

func (n *Node) String() string {
	return fmt.Sprintf(`["%s"]`, n.data)
}

func (n *Node) setNext(node *Node) {
	n.next, node.before = node, n
}

func (n *Node) Next() (node *Node, ok bool) {
	if n.next == nil {
		return n, false
	}
	return n.next, true
}

func (n *Node) Before() (node *Node, ok bool) {
	if n.before == nil {
		return n, false
	}
	return n.before, true
}

func (n *Node) getHead() (head *Node) {
	for bef, ok := n.Before(); ok; n = bef {
	}
	return n
}

func (n *Node) getTail() (tail *Node) {
	for next, ok := n.Next(); ok; n = next {
	}
	return n
}

func (n *Node) InsertAfter(node *Node) {
	next, ok := n.Next()
	n.setNext(node)
	if ok {
		node.setNext(next)
	}
}

func (n *Node) Delete() *Node {
	before, okb := n.Before()
	next, okn := n.Next()
	if okb && okn {
		before.setNext(next)
	} else if okb {
		before.next = nil
	}
	n = nil
	if okn {
		return next
	}
	return before // why choose before?
}

func (n *Node) FindExact(str string) (*Node, bool) {
	head := n.getHead()
	for {
		if head.data == str {
			return head, true
		}
		node, ok := head.Next()
		if !ok {
			break
		}
		head = node
	}
	return n, false
}

func (n *Node) FindSubStr(str string) (*Node, bool) {
	head := n.getHead()
	for {
		if strings.Contains(head.data, str) {
			return head, true
		}
		node, ok := head.Next()
		if !ok {
			break
		}
		head = node
	}
	return n, false
}

func main() {
	list := []string{"Hello", "World", "I", "am", "here", "to", "impress", "you"}
	node := NewNodes(list)
	node = node.getHead()
	fmt.Println(node)
	for {
		n, ok := node.Next()
		if ok {
			fmt.Println(n)
			node = n
			continue
		}
		fmt.Println("END")
		break
	}
	for {
		n, ok := node.Before()
		if ok {
			fmt.Println(n)
			node = n
			continue
		}
		fmt.Println("END")
		break
	}
}
