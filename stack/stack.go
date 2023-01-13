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

func (st *Stack) PrintList() {

	data := st.head

	for ele := (st.length - 1); ele >= 0; ele-- {
		fmt.Printf("[%d:%s]", ele, data.data)
		data = data.prev
	}

	fmt.Printf("\n")
}
