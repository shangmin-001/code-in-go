package sword24

import (
	. "github.com/code-in-go/algorithm/common/list"
)

// version 1.0

/**
 *
 * @param pHead ListNode类
 * @return ListNode类
 */
func ReverseList(pHead *ListNode) *ListNode {
	// write code here
	if pHead == nil {
		return pHead
	}
	pNext := pHead.Next
	if pNext != nil {
		pNewHead := ReverseList(pNext)
		pNext.Next = pHead
		pHead.Next = nil // 第一遍的时候，这个细节忽略了，造成了死循环
		return pNewHead
	}
	return pHead //没有想到优化
}

// version 2.0

func ReverseList_2(pHead *ListNode) *ListNode {
	// write code here
	if pHead == nil || pHead.Next == nil {
		return pHead
	}

	pNewHead := ReverseList_2(pHead.Next)
	pHead.Next.Next = pHead
	pHead.Next = nil
	return pNewHead
}

// version 3.0

func ReverseList_3(pHead *ListNode) *ListNode {
	// write code here

	// write code here
	var pNext *ListNode
	pCur := pHead
	var pPrev *ListNode
	for pCur != nil {
		pNext = pCur.Next
		pCur.Next = pPrev
		pPrev = pCur
		pCur = pNext
	}

	return pPrev //第一遍的时候，写的pCur，没有考虑边界值
}
