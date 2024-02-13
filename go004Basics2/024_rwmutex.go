package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Start")

	var resource Resource

	resource.name = "Resource #23"
	fmt.Println("Resource name:", resource.name)

	for i := 0; i < 10000; i++ {
		go resource.Inc()
	}

	time.Sleep(3 * time.Second)
	fmt.Println("Resource counter:", resource.Val())
}

type Resource struct {
	rwMutex sync.RWMutex
	name string
	counter int
}

func (r * Resource) Inc() {
	r.rwMutex.Lock()
	r.counter++
	r.rwMutex.Unlock()
}

func (r * Resource) Val() int {
	r.rwMutex.RLock()
	defer r.rwMutex.RUnlock()
	return r.counter
}