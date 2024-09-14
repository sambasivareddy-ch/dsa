package stack

import "fmt"

func convertPrefixToPostfixExpression(prefixExpr string) string {
	precedenceMap := make(map[byte]int)
	precedenceMap['+'] = 1
	precedenceMap['-'] = 1
	precedenceMap['*'] = 2
	precedenceMap['/'] = 2
	precedenceMap['^'] = 3

	var stack StackDS[string]
	expr_len := len(prefixExpr)
	stack.Constructor(expr_len)

	for i := expr_len - 1; i >= 0; i-- {
		chr := prefixExpr[i]
		if _, exists := precedenceMap[chr]; exists {
			operand1 := stack.Pop()
			operand2 := stack.Pop()
			expr := operand1 + operand2 + string(chr)

			stack.Push(expr)
		} else {
			stack.Push(string(chr))
		}
	}

	return stack.Pop()
}

func PrefixToPostfixExprHandler() {
	fmt.Println("*+AB-CD's postfix expression = " + convertPrefixToPostfixExpression("*+AB-CD"))
	fmt.Println("*-A/BC-/AKL's postfix expression = " + convertPrefixToPostfixExpression("*-A/BC-/AKL"))
}
