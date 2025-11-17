package customlist

import (
	"iter"
)

type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

func (list *List[T]) Push(elem T) {
	if list.head == nil {
		list.head = &element[T]{val: elem}
		list.tail = list.head
	} else {
		list.tail.next = &element[T]{val: elem}
		list.tail = list.tail.next
	}
}

func (list *List[T]) All() iter.Seq[T] {
	return func(next func(T) bool) {
		for e := list.head; e != nil; e = e.next {
			if !next(e.val) {
				return
			}
		}
	}
}

func (list *List[T]) ReverseAll() iter.Seq[T] {
	return func(next func(T) bool) {
		// collect values
		var vals []T
		for e := list.head; e != nil; e = e.next {
			vals = append(vals, e.val)
		}

		// iterate in reverse
		for i := len(vals) - 1; i >= 0; i-- {
			if !next(vals[i]) {
				return
			}
		}
	}
}
