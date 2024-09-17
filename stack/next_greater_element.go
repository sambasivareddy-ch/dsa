package stack

import "fmt"

/*
 Reference: https://www.geeksforgeeks.org/next-greater-element/

 Arr : [10, 20, 2, 5, 4, 3]
 10 -> 20
 20 -> -1
 2  ->  5
 5  -> -1
 4  -> -1
 3  -> -1
*/

func FindNextGreaterElementForEachElement(arr []int) {
	arr_len := len(arr)

	var stackDS StackDS[int]
	stackDS.Constructor(arr_len)

	for _, ele := range arr {
		if stackDS.IsEmpty() {
			stackDS.Push(ele)
			continue
		}

		for !stackDS.IsEmpty() && stackDS.Peek() < ele {
			fmt.Printf("%d => %d\n", stackDS.Pop(), ele)
		}

		stackDS.Push(ele)
	}

	for !stackDS.IsEmpty() {
		fmt.Printf("%d => %d\n", stackDS.Pop(), -1)
	}
}
