package main

import (
	"fmt"
)

const SIZE = 5

type Node struct {
	Val   string
	Left  *Node
	Right *Node
}

type Queue struct {
	Head   *Node
	Tail   *Node
	Length int
}

type Cache struct {
	Queue Queue
	Hash  Hash
}

type Hash map[string]*Node

// initial state
func NewCache() Cache {
	return Cache{Queue: NewQueue(), Hash: Hash{}}
}

// initial state
func NewQueue() Queue {
	//making a empty queue
	head := &Node{}
	tail := &Node{}

	head.Right = tail
	tail.Left = head
	return Queue{Head: head, Tail: tail}
}

func (c *Cache) Check(word string) {
	node := &Node{}
	if val, ok := c.Hash[word]; ok {
		node = c.Remove(val) //here val we get is a node
	} else {
		node = &Node{Val: word}
	}
	c.Add(node)
	c.Hash[word] = node
}
func (c *Cache) Remove(n *Node) *Node {
	fmt.Printf("Remove : %s\n", n.Val)
	left := n.Left
	right := n.Right

	left.Right = right
	right.Left = left
	c.Queue.Length -= 1
	delete(c.Hash, n.Val) //builtin function for removing a K-V pair
	return n
}

func (c *Cache) Add(n *Node) {
	fmt.Printf("Add : %s\n", n.Val)
	tmp := c.Queue.Head.Right
	//tmp here is the first element in queue other than the head

	c.Queue.Head.Right = n
	n.Left = c.Queue.Head
	n.Right = tmp
	tmp.Left = n

	c.Queue.Length++
	if c.Queue.Length > SIZE {
		fmt.Println("cache full")
		c.Remove(c.Queue.Tail.Left)
	}
}

func (c *Cache) Display() {
	c.Queue.Display()
}

func (q *Queue) Display() {
	node := q.Head.Right
	fmt.Printf("%d - [", q.Length)
	for i := 0; i < q.Length; i++ {
		fmt.Printf("{%s}", node.Val)
		if i < q.Length {
			fmt.Printf("<-->")
		}
		node = node.Right
	}
	fmt.Println("]")
}

func main() {
	fmt.Println("START, CACHE")
	cache := NewCache()
	//values to insert in cache
	for _, word := range []string{"parrot", "avocado", "tree", "potato", "tomato", "tree", "man", "monkey", "deer", "man"} {
		cache.Check(word)
		cache.Display()
	}
}
