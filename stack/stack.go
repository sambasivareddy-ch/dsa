package stack

import "fmt"

// Generic Stack
type StackDS[T any] struct {
	stack    []T
	top      int
	capacity int
}

// Initialize the capacity & top.
// If capacity == -1 then it is infinite stack (for sample 1000)
func (stackDS *StackDS[T]) Constructor(capacity int) {
	if capacity == -1 {
		stackDS.stack = make([]T, 1000)
	} else {
		stackDS.stack = make([]T, capacity)
	}
	stackDS.capacity = capacity
	stackDS.top = -1
}

func (stackDS *StackDS[T]) Push(element T) {
	if stackDS.IsFull() {
		fmt.Println("Stack Overflowed")
		return
	}
	stackDS.top++
	stackDS.stack[stackDS.top] = element
	// fmt.Printf("Inserted %v\n", element)
}

func (stackDS *StackDS[T]) IsFull() bool {
	return stackDS.top == stackDS.capacity-1
}

func (stackDS *StackDS[T]) Pop() T {
	if stackDS.IsEmpty() {
		fmt.Println("Stack is Empty!!")
		var t T
		return t
	}
	poppedEle := stackDS.stack[stackDS.top]
	stackDS.top--
	return poppedEle
}

func (stackDS *StackDS[T]) IsEmpty() bool {
	return stackDS.top == -1
}

func (stackDS *StackDS[T]) PrintStack() {
	if stackDS.IsEmpty() {
		fmt.Println("Stack is empty to print")
		return
	}

	fmt.Print("Stack: ")
	for i := 0; i <= stackDS.top; i++ {
		fmt.Printf("%v ", stackDS.stack[i])
	}
	fmt.Println()
}

func (stackDS *StackDS[T]) Peek() T {
	return stackDS.stack[stackDS.top]
}
