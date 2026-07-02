package main

import (
	"fmt"
)

type DLLNode[T any] struct {
	key   string
	value T
	prev  *DLLNode[T]
	next  *DLLNode[T]
}

type DLL[T any] struct {
	size        int
	currentSize int
	head        *DLLNode[T]
	tail        *DLLNode[T]
}

func NewDLL[T any](size int) (*DLL[T], error) {
	if size < 1 {
		return nil, fmt.Errorf("Argument \"size\" should be a positive integer. Received %d", size)
	}

	head := &DLLNode[T]{}
	tail := &DLLNode[T]{}

	head.next = tail
	tail.prev = head

	return &DLL[T]{size, 0, head, tail}, nil
}

func (dll *DLL[T]) pushHead(key string, value T) (node *DLLNode[T], deletedKey string, hasDeletion bool) {
	node = &DLLNode[T]{key: key, value: value}

	node.prev = dll.head
	node.next = dll.head.next

	dll.head.next.prev = node
	dll.head.next = node
	dll.currentSize++

	if dll.currentSize > dll.size {
		deletedKey = dll.deleteLast()
		hasDeletion = true
	}

	return
}

func (dll *DLL[T]) deleteLast() string {
	deleted := dll.tail.prev
	deletedKey := deleted.key

	prev := deleted.prev

	prev.next = dll.tail
	dll.tail.prev = prev

	deleted.prev = nil
	deleted.next = nil

	dll.currentSize--

	return deletedKey
}

func (dll *DLL[T]) deleteNode(node *DLLNode[T]) {
	node.prev.next = node.next
	node.next.prev = node.prev

	node.next = nil
	node.prev = nil

	dll.currentSize--
}
