package stack

import "fmt"

func CheckValidityOfMatchingParenthesis(parenthesisString string) bool {
	var stackDs StackDS[byte]

	stackDs.Constructor(-1) // Infinite Stack

	length := len(parenthesisString)
	for i := 0; i < length; i++ {
		parenthesis := parenthesisString[i]
		if parenthesis == '(' || parenthesis == '[' || parenthesis == '{' {
			stackDs.Push(parenthesis)
		} else {
			if stackDs.IsEmpty() {
				return false
			} else {
				peek := stackDs.Peek()
				if (parenthesis == ']' && peek == '[') ||
					(parenthesis == '}' && peek == '{') ||
					(parenthesis == ')' && peek == '(') {
					stackDs.Pop()
				} else {
					return false
				}
			}
		}
	}

	return stackDs.IsEmpty()
}

func RunParenthesisChecker() {
	fmt.Printf("()[]{}: %v\n", CheckValidityOfMatchingParenthesis("()[]{}"))
	fmt.Printf("(([]))[{}]{}: %v\n", CheckValidityOfMatchingParenthesis("(([]))[{}]{}"))
	fmt.Printf("([{[}]): %v\n", CheckValidityOfMatchingParenthesis("([{[}])"))
	fmt.Printf("(((({{{}}}))))[[[[(((()))]]]]{}: %v\n", CheckValidityOfMatchingParenthesis("(((({{{}}}))))[[[[(((()))]]]]{}"))
}
