package storage

import (
	"fmt"
	"errors"
)

type IStorage interface {
	Insert(e Employee) error
	Get(id int) (Employee, error)
	Delete(id int) error
}

type Storage struct {
	Data map[int]Employee
}

func NewStorage() *Storage {
	return &Storage {
		Data: make(map[int]Employee),
	}
}

func (s *Storage) Insert(e Employee) error {
	s.Data[e.Id] = e
	return nil
}

func (s *Storage) Get(id int) (Employee, error) {
	e, exists := s.Data[id]
	if !exists {
		err := fmt.Sprintf("Employee with id %d not found!", id)
		return Employee{}, errors.New(err)
	}
	return e, nil
}

func (s *Storage) Delete(id int) error {
	delete(s.Data, id)
	return nil
}

type Employee struct {
	Id int
	Name string
	Age int
	Salary int
}