package linked_list

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func BuildList(nodeList []int) *ListNode {
	var currentNode, firstNode *ListNode
	for _, val := range nodeList {
		listNode := &ListNode{
			Val: val,
		}

		if currentNode != nil {
			currentNode.Next = listNode
		}

		if firstNode == nil {
			firstNode = listNode
		}

		currentNode = listNode
	}

	return firstNode
}

func (listNode *ListNode) Print() {
	var stopLooping bool
	var printString string
	currentNode := listNode
	for !stopLooping {
		if currentNode != nil {
			printString += strconv.Itoa(currentNode.Val) + ","
			currentNode = currentNode.Next
		} else {
			stopLooping = true
		}
	}
	fmt.Println("hahah")
	fmt.Println(printString)
}
