package main

import (
	"fmt"
	st "StorageProject/storage"
)

func main() {
	fmt.Println("Storage Project Mod Practice")
	
	var myStorage st.IStorage
	myStorage = st.NewStorage()

	emp1 := st.Employee {
		Id: 1,
		Name: "Name1",
		Age: 1,
		Salary: 1,
	}

	myStorage.Insert(emp1)

	semp1, err := myStorage.Get(1)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(semp1.Name)

	tmpEmp, err := myStorage.Get(231)
	if err != nil {
		fmt.Println(err.Error())
	}
	
	fmt.Println(tmpEmp.Id)
}