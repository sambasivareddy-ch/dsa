package stack

import "fmt"

/*
Infix: A+B
Postfix: AB+

Infix: (A+B)*(C-D)
Postfix: AB+CD-*
*/

func ConvertInfixToPostfix(expr string) {
	precedenceMap := make(map[rune]int)
	precedenceMap['+'] = 1
	precedenceMap['-'] = 1
	precedenceMap['*'] = 2
	precedenceMap['/'] = 2
	precedenceMap['^'] = 3

	var stackDs StackDS[rune]
	stackDs.Constructor(len(expr))
	var result string = ""

	for _, chr := range expr {
		if chr == '(' {
			stackDs.Push(chr)
		} else if chr == ')' {
			for stackDs.Peek() != '(' {
				result += string(stackDs.Pop())
			}
			stackDs.Pop()
		} else if _, exists := precedenceMap[chr]; exists {
			if !stackDs.IsEmpty() {
				top := stackDs.Peek()
				for precedenceMap[top] >= precedenceMap[chr] {
					result += string(stackDs.Pop())
					top = stackDs.Peek()
				}
			}
			stackDs.Push(chr)
		} else {
			result += string(chr)
		}
	}

	for !stackDs.IsEmpty() {
		result += string(stackDs.Pop())
	}

	fmt.Println(expr + "'s postfix expression = " + result)
}
