package main

import "fmt"

func main() {
	cache, _ := NewLRU[int](2)

	cache.Update("a", 1)
	cache.Update("b", 2)

	v, ok := cache.Read("a")
	fmt.Println(v, ok) // 1 true

	cache.Update("c", 3) // b should be thrown away

	_, ok = cache.Read("b")
	fmt.Println(ok) // false

	v, ok = cache.Read("a")
	fmt.Println(v, ok) // 1 true

	v, ok = cache.Read("c")
	fmt.Println(v, ok) // 3 true
}
