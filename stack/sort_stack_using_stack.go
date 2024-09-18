package stack

/*
	Reference: https://www.geeksforgeeks.org/sort-stack-using-temporary-stack/
*/

func sortStackUsingStack(input_stack StackDS[int]) StackDS[int] {
	var temp_stack StackDS[int]

	temp_stack.Constructor(-1) // For now infinite

	for !input_stack.IsEmpty() {
		top_ele := input_stack.Pop()

		// while temporary stack is not empty and top
		// of stack is lesser than temp
		for !temp_stack.IsEmpty() && temp_stack.Peek() < top_ele {
			input_stack.Push(temp_stack.Pop())
		}

		temp_stack.Push(top_ele)
	}

	return temp_stack
}

func DemonstrateStackSortUsingStack() {
	var input_stack StackDS[int]
	input_stack.Constructor(-1)
	input_stack.Push(100)
	input_stack.Push(30)
	input_stack.Push(35)
	input_stack.Push(20)
	input_stack.Push(120)
	result_stack := sortStackUsingStack(input_stack)
	result_stack.PrintStack()

	var input_stack1 StackDS[int]
	input_stack1.Constructor(-1)
	input_stack1.Push(-1)
	input_stack1.Push(0)
	input_stack1.Push(2)
	input_stack1.Push(200)
	input_stack1.Push(1230)
	input_stack1.Push(340)
	result_stack1 := sortStackUsingStack(input_stack1)
	result_stack1.PrintStack()
}
