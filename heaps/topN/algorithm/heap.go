package algorithm

import (
	"fmt"
	"math"
)

type Node struct {
	value  int
	left   *Node
	right  *Node
	parent *Node
}

func New() *Node {
	return &Node{
		value: -1,
	}
}

func (n *Node) GetMin() int {
	return n.value
}

func (n *Node) ExtractMin() int {

	minimumValue := n.value

	if n.parent == nil && n.left == nil && n.right == nil {
		*n = *New()
		return minimumValue
	}

	var q queue
	q.enqueue(n)
	var previous *Node
	for !q.isEmpty() {
		current := q.dequeue()

		if current.left == nil {
			n.value = previous.right.value
			previous.right = nil
			n.balanceWithChildren()
			break
		} else if current.right == nil {
			n.value = current.left.value
			current.left = nil
			n.balanceWithChildren()
			break
		}
		q.enqueue(current.left)
		q.enqueue(current.right)
		previous = current
	}

	return minimumValue
}

func (n *Node) Insert(value int) {

	// when heap is empty
	if n.value == -1 {
		n.value = value
		return
	}

	var q queue
	q.enqueue(n)

	for !q.isEmpty() {
		current := q.dequeue()
		if current.left == nil {
			current.left = &Node{
				parent: current,
				value:  value,
			}
			current.left.balanceWithParent()
			break
		} else if current.right == nil {
			current.right = &Node{
				parent: current,
				value:  value,
			}
			current.right.balanceWithParent()
			break
		}
		q.enqueue(current.left)
		q.enqueue(current.right)
	}
}

func (n *Node) balanceWithParent() {
	if n.parent == nil {
		return
	}
	if n.value < n.parent.value {
		n.value, n.parent.value = n.parent.value, n.value
		n.parent.balanceWithParent()
	} else {
		return
	}
}

func (n *Node) balanceWithChildren() {
	if n.left == nil {
		return
	} else if n.right == nil {
		if n.value > n.left.value {
			n.value, n.left.value = n.left.value, n.value
		}
	} else {
		minNode := getMinimumValueNode(n.left, n.right)
		if n.value > minNode.value {
			n.value, minNode.value = minNode.value, n.value
			minNode.balanceWithChildren()
		}
	}
}

func (n *Node) PrintHeap() {
	if n == nil {
		return
	}
	fmt.Printf("node: %v ", n.value)
	if n.left != nil {
		fmt.Printf("left: %v ", n.left.value)
	} else {
		fmt.Printf("left: nil ")
	}
	if n.right != nil {
		fmt.Printf("right: %v ", n.right.value)
	} else {
		fmt.Printf("right: nil ")
	}
	fmt.Println()

	n.left.PrintHeap()
	n.right.PrintHeap()
}

func getMinimumValueNode(node1, node2 *Node) *Node {
	switch math.Min(float64(node1.value), float64(node2.value)) {
	case float64(node1.value):
		return node1
	case float64(node2.value):
		return node2
	}
	return nil
}
