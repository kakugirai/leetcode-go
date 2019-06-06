package main

import (
	"container/list"
	"fmt"
)

type LRUCache struct {
	cap int                   // capacity
	l   *list.List            // doubly linked list
	m   map[int]*list.Element // hash table for checking if list node exists
}

// Pair is the value of a list node.
type Pair struct {
	key   int
	value int
}

// Constructor initializes a new LRUCache.
func Constructor(capacity int) LRUCache {
	return LRUCache{
		cap: capacity,
		l:   new(list.List),
		m:   make(map[int]*list.Element, capacity),
	}
}

// Get a list node from the hash map.
func (c *LRUCache) Get(key int) int {
	// check if list node exists
	if node, ok := c.m[key]; ok {
		val := node.Value.(*list.Element).Value.(Pair).value
		// move node to front
		c.l.MoveToFront(node)
		return val
	}
	return -1
}

// Put key and value in the LRUCache
func (c *LRUCache) Put(key int, value int) {
	// check if list node exists
	if node, ok := c.m[key]; ok {
		// move the node to front
		c.l.MoveToFront(node)
		// update the value of a list node
		node.Value.(*list.Element).Value = Pair{key: key, value: value}
	} else {
		// delete the last list node if the list is full
		if c.l.Len() == c.cap {
			// get the key that we want to delete
			idx := c.l.Back().Value.(*list.Element).Value.(Pair).key
			// delete the node pointer in the hash map by key
			delete(c.m, idx)
			// remove the last list node
			c.l.Remove(c.l.Back())
		}
		// initialize a list node
		node := &list.Element{
			Value: Pair{
				key:   key,
				value: value,
			},
		}
		// push the new list node into the list
		ptr := c.l.PushFront(node)
		// save the node pointer in the hash map
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
