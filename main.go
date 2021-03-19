package main

import (
	"fmt"
	"golang-queue-list/queue"
	"golang-queue-list/stack"
)

func main() {
	stackImpl := new(stack.Stack)
	stackImpl.Push(1)
	stackImpl.Push(2)
	if val, ok := stackImpl.Peek(); ok {
		fmt.Printf("Stack Peek value %v current size: %v\n", val, stackImpl.Size())
	}
	if val, ok := stackImpl.Pop(); ok {
		fmt.Printf("Stack Popped value %v current size: %v\n", val, stackImpl.Size())
	}
	if val, ok := stackImpl.Peek(); ok {
		fmt.Printf("Stack Peek value %v current size: %v\n", val, stackImpl.Size())
	}
	if val, ok := stackImpl.Pop(); ok {
		fmt.Printf("Stack Popped value %v current size: %v\n", val, stackImpl.Size())
	}

	queueImpl := new(queue.Queue)

	queueImpl.Push(1)
	queueImpl.Push(2)
	if val, ok := queueImpl.Peek(); ok {
		fmt.Printf("Queue Peek value %v current size: %v\n", val, queueImpl.Size())
	}

	if val, ok := queueImpl.Pop(); ok {
		fmt.Printf("Queue Pop value %v current size: %v\n", val, queueImpl.Size())
	}

	if val, ok := queueImpl.Peek(); ok {
		fmt.Printf("Queue Peek value %v current size: %v\n", val, queueImpl.Size())
	}

	if val, ok := queueImpl.Pop(); ok {
		fmt.Printf("Queue Pop value %v current size: %v\n", val, queueImpl.Size())
	}
}
