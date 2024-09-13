package stack

import "fmt"

type QueueWithStack[T any] struct {
	stack1 StackDS[T]
	stack2 StackDS[T]
}

func (qws *QueueWithStack[T]) constructQueue() {
	// Initialize with Maximum Capacity
	qws.stack1.Constructor(-1)
	qws.stack2.Constructor(-1)
}

func (qws *QueueWithStack[T]) enqueue(elem T) {
	for !qws.stack1.IsEmpty() {
		qws.stack2.Push(qws.stack1.Pop())
	}

	qws.stack1.Push(elem)

	for !qws.stack2.IsEmpty() {
		qws.stack1.Push(qws.stack2.Pop())
	}

	fmt.Printf("Inserted %v\n", elem)
}

func (qws *QueueWithStack[T]) dequeue() T {
	for qws.stack1.IsEmpty() {
		fmt.Println("No elements exists to dequeue")
		var t T
		return t
	}
	return qws.stack1.Pop()
}

func (qws *QueueWithStack[T]) printQueue() {
	qws.stack1.PrintStack()
}

func ExecuteQueueWithStack() {
	var qws QueueWithStack[int]

	qws.constructQueue() // For now with maxium capacity

	for 1%10 == 1 {
		var option int
		fmt.Println("Options:")
		fmt.Println("1. Enqueue")
		fmt.Println("2. Dequeue")
		fmt.Println("3. Print Queue")
		fmt.Println("4. Break")
		fmt.Print("Enter Option: ")
		fmt.Scan(&option)

		if option == 1 {
			var elem int
			fmt.Print("Enter element: ")
			fmt.Scan(&elem)
			qws.enqueue(elem)
		} else if option == 2 {
			elem := qws.dequeue()
			fmt.Printf("Removed %v\n", elem)
		} else if option == 3 {
			qws.printQueue()
		} else {
			break
		}
	}
}
