package sword24

import "testing"
import "github.com/code-in-go/algorithm/common/list"

func TestReverseList_2(t *testing.T) {
	head := list.NewLinkedList([]int{1, 2, 3, 4, 5})
	list.PrintLinkedList(head)
	head = ReverseList(head)
	list.PrintLinkedList(head)
}
