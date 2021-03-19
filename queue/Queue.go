package queue

import (
	"sync"
)

type Queue struct {
	queue []int
	lock  sync.RWMutex
}

func (c *Queue) New() *Queue {
	return &Queue{
		queue: make([]int, 0),
	}
}

func (c *Queue) Push(val int) bool {
	c.lock.Lock()
	defer c.lock.Unlock()
	currentLen := len(c.queue)
	c.queue = append(c.queue, val)
	return currentLen < len(c.queue)
}

func (c *Queue) Pop() (int, bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	if len(c.queue) == 0 {
		return -1, false
	}

	popVal := c.queue[0]
	c.queue = c.queue[1:]
	return popVal, true
}

func (c *Queue) Peek() (int, bool) {
	c.lock.Lock()
	defer c.lock.Unlock()
	if len(c.queue) == 0 {
		return -1, false
	}
	peekVal := c.queue[0]
	return peekVal, true
}

func (c *Queue) Empty() bool {
	c.lock.Lock()
	defer c.lock.Unlock()
	return len(c.queue) == 0
}

func (c *Queue) Size() int {
	c.lock.Lock()
	defer c.lock.Unlock()
	return len(c.queue)
}
