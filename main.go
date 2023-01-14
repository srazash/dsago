package main

import (
	"fmt"
	"time"

	//"github.com/srazash/dsago/queue"
	"github.com/srazash/dsago/quicksort"
	//"github.com/srazash/dsago/stack"
)

func main() {

	start := time.Now()

	/* myQueue := queue.Queue{}

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

	fmt.Printf("Is the queue valid? %t (%d:%d)", valid, count, len) */

	/* myStack := stack.Stack{}

	fmt.Printf("Stack length: %d\n", myStack.Length())

	myStack.Push("This")
	myStack.Push("is")
	myStack.Push("a")
	myStack.Push("stack!")

	myStack.PrintList()

	fmt.Printf("Stack length: %d\n", myStack.Length())

	fmt.Printf("POP!: %s\n", myStack.Pop())
	fmt.Printf("POP!: %s\n", myStack.Pop())

	fmt.Printf("Stack length: %d\n", myStack.Length())

	fmt.Printf("Peak: %s\n", myStack.Peek())

	myStack.PrintList() */

	array := []int{7, 9, 11, 4, 1, 2, 12, 6}

	fmt.Println(array)

	quicksort.QuickSort(array, 0, len(array)-1)

	fmt.Println(array)

	fmt.Printf("runtime: %dns\n", time.Since(start).Nanoseconds())

}
