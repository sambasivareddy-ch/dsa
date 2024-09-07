package stack

import "fmt"

func ReverseString(str string) {
	var stack StackDS[rune]

	stack.Constructor(len(str))

	for _, j := range str {
		stack.Push(j)
	}

	resultStr := ""
	for !stack.IsEmpty() {
		resultStr += string(stack.Pop())
	}

	fmt.Println(str + " reversed result: " + resultStr)
}
