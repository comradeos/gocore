package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Start")
	var resource Resource
	resource.Counter = make(map[string]int)
	resource.Name = "Resource #1"

	fmt.Println("Resource name:", resource.Name)

	for i := 0; i < 10000; i++ {
		go resource.IncrementCounter("key1")
	}

	time.Sleep(3 * time.Second)
	fmt.Println("Resource counter:", resource.GetCounter("key1"))
}

type Resource struct {
	Mutex sync.Mutex
	Name string
	Counter map[string]int
}

func (r * Resource) IncrementCounter(key string) {
	r.Mutex.Lock()
	r.Counter[key]++
	r.Mutex.Unlock()
}

func (r * Resource) DecrementCounter(key string) {
	r.Mutex.Lock()
	r.Counter[key]--
	r.Mutex.Unlock()
}

func (r * Resource) GetCounter(key string) int {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()
	return r.Counter[key]
}