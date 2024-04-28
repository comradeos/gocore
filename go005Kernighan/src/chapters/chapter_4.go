package chapters

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

func DataTypes2() {
	fmt.Println("DataTypes2")

	a := []byte("Hello")
	b := []byte("Hello")

	sha1 := sha256.Sum256(a)
	sha2 := sha256.Sum256(b)

	fmt.Printf("%x\n", sha1)
	fmt.Printf("%x\n", sha2)
}

func ArraysSlicesMaps() {
	myArray := [3]int{1, 2, 3}
	fmt.Println(myArray)

	mySlice := []int{1, 2, 3}
	fmt.Println(mySlice)
	mySlice = append(mySlice, 4)
	fmt.Println(mySlice)

	myMap := map[string]int{}
	myMap["one"] = 1
	myMap["two"] = 2
	myMap["three"] = 3

	fmt.Println(myMap)

	for key, value := range myMap {
		fmt.Println(key, value)
	}

	fmt.Println("-------------------------")

	type Struct1 struct{}

	var s0 Struct1
	s0 = Struct1{}
	fmt.Println(s0) // {}

	var s1 = Struct1{}
	fmt.Println(s1) // {}

	s2 := Struct1{}
	fmt.Println(s2) // {}

	s3 := new(Struct1)
	fmt.Println(s3) // &{}

	var s4 *Struct1
	fmt.Println(s4) // <nil>
}

func showCap(s *[]int) {
	*s = append(*s, (*s)[len(*s)-1]+1)
	fmt.Println(*s, cap(*s))
}

func LenCap() {
	var mySlice = []int{1}
	fmt.Println(mySlice, cap(mySlice))

	for range 20 {
		showCap(&mySlice)
	}
}

func Structs() {
	type Employee struct {
		ID        int
		Name      string
		Address   string
		DoB       time.Time // Date of Birth
		Position  string
		Salary    int
		ManagerID int
	}

	var employee = Employee{
		ID: 1,
	}

	var pEmployee = &employee
	var ppEmployee = &pEmployee

	(*pEmployee).Name = "John"

	fmt.Println((**ppEmployee).Name)

}

func Trees() {
	type tree struct {
		value       int
		left, right *tree
	}
	var t1 = tree{value: 1}
	t1.right = new(tree)
	t1.right.left = new(tree)
	t1.right.left.value = 555

	fmt.Println(t1.right.left.value)
}

func Structs2() {
	type Point struct {
		X int
		Y int
		Z int
	}

	var p1 = Point{
		X: 1,
		Y: 2,
		Z: 3,
	}

	fmt.Println(p1)

	var a = Point{X: 1, Y: 2, Z: 3}
	var b = Point{X: 2, Y: 2, Z: 3}
	fmt.Println(a == b)

	fmt.Println("-------------------------")

	type Circle struct {
		X, Y, Radius int
	}

	type Box struct {
		Width, Height, Depth int
	}

	type Wheel struct {
		Circle
		Box
		Weight int
	}

	var w1 = Wheel{}
	fmt.Println(w1)
	w1.X = 8
	w1.Y = 8
	w1.Radius = 8
	w1.Weight = 8
	w1.Width = 8
	w1.Height = 8
	w1.Depth = 8

	fmt.Println(w1)

	var c1 = Circle{X: 1, Y: 2, Radius: 3}
	fmt.Println(c1)
}

func Jsons() {
	type Person struct {
		Name string
		Age  int
	}

	var p1 = Person{
		Name: "John",
		Age:  30,
	}

	jsonP1, err := json.Marshal(p1)
	if err != nil {
		fmt.Println(err)
	}

	jsonP1, err = json.MarshalIndent(p1, "", "  ")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(jsonP1))

	type Object struct {
		Id    int
		Value string
	}

	var jsonStr = `{"Id": 1, "Value": "Hello"}`
	var jsonBytes = []byte(jsonStr)

	var obj Object
	err = json.Unmarshal(jsonBytes, &obj)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(obj)

	log.Fatalf("Error: %v", "aaaaaa")
}

func Templates() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(dir)

	file, err := os.Open("files/my_template.txt")
	if err != nil {
		log.Fatal(err)
	}

	fileContent := make([]byte, 100)
	_, err = file.Read(fileContent)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(fileContent))

}
