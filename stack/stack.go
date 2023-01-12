package stack

import (
	"fmt"
	"log"
)

type element struct {
	prev *element
	data string
}

type Stack struct {
	head   *element
	length int
}

func (st *Stack) Push(s string) {

	ele := element{}
	ele = element{prev: &ele, data: s}

	if st.head == nil {
		st.head = &ele
	}

	ele.prev = st.head
	st.head = &ele

	st.length++
}

func (st *Stack) Pop() string {

	if st.length == 0 {
		log.Fatal("unable to pop: stack is empty")
		return ""
	}

	st.length--

	if st.length == 0 {
		hd := st.head.data
		st.head = nil
		return hd
	}

	hd := st.head.data
	st.head = st.head.prev
	return hd
}

func (st *Stack) Peek() string {
	return st.head.data
}

func (st *Stack) Length() int {
	return st.length
}

func (st *Stack) PrintStack() {
	ele := st.head
	count := 0

	for count < st.length {
		fmt.Printf("[%d: %s]", count, ele.data)

		if ele != st.head {
			fmt.Printf("--->")
		} else {
			fmt.Printf("(h)\n")
		}

		ele = ele.prev
		count++
	}

}

func (st *Stack) ValidStack() (bool, int, int) {
	element := st.head
	count := 0
	len := st.length

	for count < len {
		element = element.prev
		count++
	}

	if count == len {
		return true, count, len
	} else {
		return false, count, len
	}

}

func (st *Stack) Head() *element {
	return st.head
}
