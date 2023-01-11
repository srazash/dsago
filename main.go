package main

import (
	"fmt"

	"github.com/srazash/dsago/queue"
)

func main() {

	str := "Hellorld!"
	fmt.Printf("%s\n", str)

	myQueue := queue.Queue{}
	fmt.Printf("Queue length is %d [h: %v t: %v]\n", myQueue.Length(), myQueue.Head(), myQueue.Tail())

	myQueue.Enqueue("Hello")
	myQueue.Enqueue("Ben")
	myQueue.Enqueue("how")
	myQueue.Enqueue("are")
	myQueue.Enqueue("you?")

	fmt.Printf("Queue length is %d [h: %v t: %v]\n", myQueue.Length(), myQueue.Head(), myQueue.Tail())
	myQueue.PrintQueue()

	myQueue.Dequeue()
	fmt.Printf("Queue length is %d [h: %v t: %v]\n", myQueue.Length(), myQueue.Head(), myQueue.Tail())
	myQueue.PrintQueue()

}
