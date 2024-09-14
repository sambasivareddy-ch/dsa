# Stack Applications

### Reverse String
1. Read the input string
2. Read one by one character in the string from left to right and push each character to the top of stack
3. Initialize a empty string
4. Then read each character from the stack by popping the top character at a time, and appended to the result string
5. Return the result string
<hr>

### Parenthesis Checker
1. Intialize an empty character stack.
2. Push all the open brackets like '[', '{', '(' to the top of the stack.
3. If any close bracket encounters like ']', '}', ')', then check whether top of the stack is it's respective open bracket. If yes continue until processing all brackets else return false.
<hr>

### Queue With Stack
1. A queue (follows FIFO) can be implemented using two stacks, for say stack1 & stack2.
#### Enqueue
1. Enqueue is the operation to insert the element to the front of the queue.
2. To implement this, we can make use of two stacks: stack1 & stack2
3. When we are inserting the new element to queue, pop all the elements from stack1 & push that element to stack2 simultaneously.
4. After pushing stack1 elements to stack2, push the "new element" to the stack1
5. Then pop all the elements from stack2 & push them to stack1 simultaneously
#### Dequeue
1. Dequeue is the operation to delete the element from the rear side of the queue
2. To implement this, we can simple pop the top element of the stack1 and return that popped element
<hr>

