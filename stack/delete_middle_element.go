package stack

/*
	Reference: https://www.geeksforgeeks.org/delete-middle-element-stack/
*/

func DeleteMiddleElementFromStack[T any](stack *StackDS[T], current int, size int) {
	if current == size/2 {
		stack.Pop()
		return
	}

	top := stack.Pop()
	DeleteMiddleElementFromStack(stack, current+1, size)

	// Append popped elements
	stack.Push(top)
}

/*
	package main

	import (
		"github.com/sambasivareddy-ch/dsa/stack"
	)

	func main() {
		var stackDS stack.StackDS[int]

		stackDS.Constructor(7)

		stackDS.Push(10)
		stackDS.Push(20)
		stackDS.Push(30)
		stackDS.Push(40)
		stackDS.Push(50)
		stackDS.Push(60)
		stackDS.Push(70)

		stack.DeleteMiddleElementFromStack(&stackDS, 0, stackDS.Size())

		stackDS.PrintStack() // 10 20 30 50 60 70
	}

*/
