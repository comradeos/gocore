package main

import f"fmt"
import h"go003Basics/helper"
import "errors"

type employee struct {
	id int
	name string
	age int
	salary int
}

type storage interface {
	insert(e employee) error
	get(id int) (employee error)
	delete(id int) error
}

type memoryStorage struct {
	data map[int]employee
}

func newMemoryStorage() *memoryStorage {
	return &memoryStorage {
		data: make(map[int]employee),
	}
}

func (ms *memoryStorage) insert(e employee) error {
	ms.data[e.id] = e
	return nil
}

func (ms *memoryStorage) get(id int) (employee, error) {
	e, exists := ms.data[id]
	if !exists {
		err := f.Sprintf("employee with id %d not found", id)
		return employee{}, errors.New(err)
	}
	return e, nil
}

func (ms *memoryStorage) delete(id int) error {
	delete(ms.data, id)
	return nil
}

func main() {
	f.Println("Interfaces:")
	h.Sep()
	
	var storage = newMemoryStorage()
	
	emp1 := employee {
		id: 1,
		name: "One",
		age: 22,
		salary: 1000,
	}	

	emp2 := employee {
		id: 2,
		name: "One",
		age: 32,
		salary: 2000,
	}

	storage.insert(emp1)
	storage.insert(emp2)

	tmpEmp, err := storage.get(1)

	if err != nil {
		f.Println(err)
	} else {
		f.Println(tmpEmp)
	}

	h.Sep()

	err = storage.delete(123)
	if err != nil {
		f.Println(err)
	}
	
	f.Println(storage)
}

