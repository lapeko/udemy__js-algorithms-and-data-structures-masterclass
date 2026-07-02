package main

import "fmt"

type LRU[T any] interface {
	Read(string) (T, bool)
	Update(string, T)
}

type lru[T any] struct {
	size  int
	table map[string]*DLLNode[T]
	dll   *DLL[T]
}

func NewLRU[T any](size int) (LRU[T], error) {
	if size < 1 {
		return nil, fmt.Errorf("Argument \"size\" should be a positive integer. Received %d", size)
	}
	dll, err := NewDLL[T](size)
	if err != nil {
		return nil, err
	}
	return &lru[T]{
		size,
		make(map[string]*DLLNode[T]),
		dll,
	}, nil
}

func (lru *lru[T]) Read(s string) (T, bool) {
	node, ok := lru.table[s]

	if !ok {
		var value T
		return value, ok
	}

	lru.dll.deleteNode(node)
	node, _, _ = lru.dll.pushHead(s, node.value)
	lru.table[s] = node

	return node.value, ok
}

func (lru *lru[T]) Update(s string, value T) {
	oldNode, ok := lru.table[s]
	if ok {
		lru.dll.deleteNode(oldNode)
	}
	node, deletedKey, hasDeletedTail := lru.dll.pushHead(s, value)
	if hasDeletedTail {
		delete(lru.table, deletedKey)
	}
	lru.table[s] = node
}
