package main

import (
	"container/list"
	"fmt"
)

type LRUCache struct {
	cap int
	l   *list.List
	m   map[int]*list.Element
}

type Pair struct {
	index int
	value int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cap: capacity,
		l:   new(list.List),
		m:   make(map[int]*list.Element, capacity),
	}
}

func (c *LRUCache) Get(key int) int {
	if node, ok := c.m[key]; ok {
		res := node.Value.(*list.Element).Value.(Pair).value
		c.l.MoveToFront(node)
		return res
	}
	return -1
}

func (c *LRUCache) Put(key int, value int) {
	if node, ok := c.m[key]; ok {
		c.l.MoveToFront(node)
		node.Value.(*list.Element).Value = Pair{index: key, value: value}
	} else {
		if c.l.Len() == c.cap {
			ele := c.l.Back().Value.(*list.Element).Value.(Pair)
			delete(c.m, ele.index)
			c.l.Remove(c.l.Back())
		}
		node := &list.Element{
			Value: Pair{
				index: key,
				value: value,
			},
		}
		ptr := c.l.PushFront(node)
		c.m[key] = ptr
	}
}

func main() {
	obj := Constructor(2)   // nil
	obj.Put(1, 10)          // nil, linked list: [1:10]
	obj.Put(2, 20)          // nil, linked list: [2:20, 1:10]
	fmt.Println(obj.Get(1)) // 10, linked list: [1:10, 2:20]
	obj.Put(3, 30)          // nil, linked list: [1:10, 3:30]
	fmt.Println(obj.Get(2)) // should be -1, but i got 20
	obj.Put(4, 40)          // nil, linked list: [4:40, 1:10]
	fmt.Println(obj.Get(1)) // 10, linked list: [1:10, 4:40]
	fmt.Println(obj.Get(3)) // -1, linked list: [1:10, 4:40]
	fmt.Println(obj.Get(4)) // 40, linked list: [4:40, 1:10]
}
