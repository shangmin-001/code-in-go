package sword6

import . "github.com/code-in-go/algorithm/common/list"

func printListFromTailToHead_1(head *ListNode) []int {
	// write code here
	res := make([]int, 0)
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	for i := 0; i < len(res)-1-i; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-1-i], res[i]
	}
	return res
}

func printListFromTailToHead(head *ListNode) []int {
	// write code here
	if head == nil {
		return nil
	}

	res := printListFromTailToHead(head.Next)
	if res != nil {
		return append(res, head.Val)
	}
	return []int{head.Val}

}
