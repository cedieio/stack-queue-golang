package stack

import (
	"sync"
)

type Stack struct {
	stack []int
	lock  sync.RWMutex
}

func (c *Stack) New() *Stack {
	return &Stack{
		stack: make([]int, 0),
	}
}

func (c *Stack) Push(val int) bool {
	c.lock.Lock()
	defer c.lock.Unlock()
	currentLen := len(c.stack)
	c.stack = append(c.stack, val)
	return currentLen < len(c.stack)
}

func (c *Stack) Pop() (int, bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	if len(c.stack) == 0 {
		return -1, false
	}
	lastIndex := len(c.stack) - 1
	popVal := c.stack[lastIndex]
	c.stack = c.stack[:lastIndex]
	return popVal, true
}

func (c *Stack) Peek() (int, bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	if len(c.stack) == 0 {
		return -1, false
	}
	lastIndex := len(c.stack) - 1
	popVal := c.stack[lastIndex]
	return popVal, true
}

func (c *Stack) Empty() bool {
	c.lock.Lock()
	defer c.lock.Unlock()
	return len(c.stack) == 0
}

func (c *Stack) Size() int {
	c.lock.Lock()
	defer c.lock.Unlock()
	return len(c.stack)
}
