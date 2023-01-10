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

	e := element{}
	e = element{next: &e, data: s}

	if q.head == nil {
		q.head, q.tail = &e, &e
	}

	q.tail.next, q.tail = &e, &e

	q.length++
}

func (q *Queue) Dequeue() {

	if q.head == nil {
		log.Fatal("unable to dequeue: queue head is nil")
	} else {
		q.head = q.head.next
		q.length--
	}

}

func (q *Queue) Length() int {
	return q.length
}

func (q *Queue) Peak(e int) string {
	return ""
}

func (q *Queue) PrintQueue() {
	e := q.head
	c := 0

	for c < q.length {
		fmt.Printf("[%d: %s]", c, e.data)

		if e != q.tail {
			fmt.Printf("->")
		} else {
			fmt.Printf("\n")
		}

		e = e.next
		c++
	}

}

func (q *Queue) Head() *element {
	return q.head
}

func (q *Queue) Tail() *element {
	return q.tail
}
