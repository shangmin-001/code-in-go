package sword9

var stack1 []int
var stack2 []int

// 第一次解题的时候

func Push_1(node int) {
	stack1 = append(stack1, node)
}

func Pop_2() int {
	res := 0
	if len(stack2) != 0 {
		res = popStack2(&stack2)
		return res
	}

	if len(stack1) == 0 {
		return res
	}

	for i := len(stack1) - 1; i >= 0; i-- {
		stack2 = append(stack2, stack1[i])
	}
	stack1 = nil
	if len(stack2) != 0 {
		res = popStack2(&stack2)
	}
	return res

}

func popStack2(stack2 *[]int) int {
	res := 0

	res = (*stack2)[len(*stack2)-1]
	if len(*stack2) == 1 {
		*stack2 = nil
	} else {
		*stack2 = (*stack2)[0 : len(*stack2)-1]
	}
	return res

}

// 看了别人的解法时候重新重写了一遍

func Push(node int) {
	stack1 = append(stack1, node)
}

func Pop() int {
	if len(stack2) == 0 {
		for i := len(stack1) - 1; i >= 0; i-- {
			stack2 = append(stack2, stack1[i])
		}
		stack1 = stack1[:0]
	}

	res := stack2[len(stack2)-1]
	stack2 = stack2[:len(stack2)-1]
	return res
}
