package list

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewLinkedList(data []int) (pHead *ListNode) {

	var pPrev *ListNode
	for _, item := range data {
		if pHead == nil {
			pHead = NewListNode(item)
			pPrev = pHead
			continue
		}
		pPrev.Next = NewListNode(item)
		pPrev = pPrev.Next
	}
	return pHead
}

func PrintLinkedList(pHead *ListNode) {
	if pHead == nil {
		fmt.Println("")
		return
	}
	fmt.Print("->", pHead.Val)
	PrintLinkedList(pHead.Next)
}

func NewListNode(val int) *ListNode {
	return &ListNode{Val: val}
}
