package sword25

import . "github.com/code-in-go/algorithm/common/list"

func Merge(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	// write code here
	if pHead1 == nil {
		return pHead2
	}
	if pHead2 == nil {
		return pHead1
	}

	if pHead1.Val < pHead2.Val {
		pHead1.Next = Merge(pHead1.Next, pHead2)
		return pHead1
	}
	pHead2.Next = Merge(pHead1, pHead2.Next)
	return pHead2
}

// 优化前
func Merge_1(pHead1 *ListNode, pHead2 *ListNode) *ListNode {
	// write code here
	if pHead1 == nil {
		return pHead2
	}
	if pHead2 == nil {
		return pHead1
	}
	var subPHead *ListNode
	var res *ListNode
	if pHead1.Val < pHead2.Val {
		res = pHead1
		subPHead = Merge_1(pHead1.Next, pHead2)
	} else {
		res = pHead2
		subPHead = Merge_1(pHead1, pHead2.Next)
	}
	res.Next = subPHead
	return res
}

//非递归也不写了
