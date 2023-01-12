package main

import (
	"fmt"

	"github.com/srazash/dsago/queue"
	"github.com/srazash/dsago/stack"
)

func main() {

	myQueue := queue.Queue{}

	myQueue.Enqueue("Hello")
	myQueue.Enqueue("Ben")
	myQueue.Enqueue("how")
	myQueue.Enqueue("are")
	myQueue.Enqueue("you?")

	fmt.Printf("Queue length: %d\n", myQueue.Length())

	myQueue.PrintQueue()

	myQueue.Dequeue()

	myQueue.PrintQueue()

	myQueue.Enqueue("Hello")

	myQueue.PrintQueue()

	valid, count, len := myQueue.ValidQueue()

	fmt.Printf("Is the queue valid? %t (%d:%d)", valid, count, len)

	myStack := stack.Stack{}

	myStack.Push("This")
	myStack.Push("is")
	myStack.Push("a")
	myStack.Push("stack!")

	myStack.PrintStack()

	_ = myStack.Pop()
	_ = myStack.Pop()

	myStack.PrintStack()

}
