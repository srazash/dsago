package queue

import (
	"fmt"
	"log"
)

type element struct {
	next *element
	data string
}

type Queue struct {
	head   *element
	tail   *element
	length int
}

func (q *Queue) Enqueue(s string) {

	ele := element{}
	ele = element{next: &ele, data: s}

	if q.head == nil {
		q.head, q.tail = &ele, &ele
	}

	q.tail.next, q.tail = &ele, &ele

	q.length++
}

func (q *Queue) Dequeue() string {

	if q.head == nil {
		log.Fatal("unable to dequeue: queue is empty")
		return ""
	}

	q.length--

	head := q.head.data
	q.head = q.head.next
	return head

}

func (q *Queue) Peak() string {
	return q.head.data
}

func (q *Queue) Length() int {
	return q.length
}

func (q *Queue) PrintQueue() {
	ele := q.head
	count := 0

	fmt.Printf("(h)")

	for count < q.length {
		fmt.Printf("[%d: %s]", count, ele.data)

		if ele != q.tail {
			fmt.Printf("<---")
		} else {
			fmt.Printf("(t)\n")
		}

		ele = ele.next
		count++
	}

}

func (q *Queue) ValidQueue() (bool, int, int) {
	element := q.head
	count := 0
	len := q.Length()

	for count < len {
		element = element.next
		count++
	}

	if count == len {
		return true, count, len
	} else {
		return false, count, len
	}

}

func (q *Queue) Head() *element {
	return q.head
}

func (q *Queue) Tail() *element {
	return q.tail
}
