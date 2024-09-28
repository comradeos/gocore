package chapters

import (
	"fmt"
)

func Ch004Functions() {
	fmt.Println("============== Hello from Ch004Functions() ==============")

	// функция суммирования
	fmt.Println(Sum(1, 2))

	// возврат нескольких значений
	sum, sub := SumAndSub(10, 5)
	fmt.Println("sum", sum, "sub", sub)

	// возврат нескольких значений с ошибкой
	sum, sub, err := SumAndSubWithError(10, 5)
	if err != nil {
		fmt.Println("error", err)
	} else {
		fmt.Println("sum", sum, "sub", sub)
	}

	sum, sub, err = SumAndSubWithError(5, 10)
	if err != nil {
		fmt.Println("error", err)
	} else {
		fmt.Println("sum", sum, "sub", sub)
	}

	// суммирование произвольного количества аргументов
	fmt.Println(SumVariadic(1, 2, 3, 4, 5))

	var slice1 = []int{1, 2, 3, 4, 5}
	fmt.Println(SumVariadic(slice1...))

	// анонимные функции
	func(p string) {
		fmt.Println("Hello from anonymous function: " + p)
	}("myParam")

	// анонимные функции как переменные
	var myFunc = func(p string) {
		fmt.Println("Hello from anonymous function: " + p)
	}
	myFunc("myParam")

	// функция как тип
	var add MyFuncType = Sum
	result := add(1, 2)
	fmt.Println("result", result)

	// функция обратного вызова
	var myCallBack MyTypeForCallBack = func(a int, b int) {
		fmt.Println("Hello from myCallBack function"+" a: ", a, " b: ", b)
	}

	FuncWithCallBack(1, 2, myCallBack)

	// замыкания
	prefixer := func(prefix string) Prefixer {
		return func(value string) string {
			return prefix + value
		}
	}

	prefix := prefixer("prefix: ")
	fmt.Println(prefix("value"))

	// defer
	//DeferExample()

	// panic
	//PanicExample()

	// recover
	//RecoverExample()
	//RecoverExample2()

	DeferPanicsTest()

}

// Sum функция суммирования
func Sum(a int, b int) int {
	return a + b
}

// SumAndSub Возврат нескольких значений
func SumAndSub(a int, b int) (int, int) {
	return a + b, a - b
}

// SumAndSubWithError Возврат нескольких значений с ошибкой
func SumAndSubWithError(a int, b int) (int, int, error) {
	if a < b {
		return 0, 0, fmt.Errorf("a < b")
	}

	return a + b, a - b, nil
}

// SumVariadic Суммирование произвольного количества аргументов
func SumVariadic(args ...int) int {
	sum := 0
	for _, arg := range args {
		sum += arg
	}
	return sum
}

// MyFuncType функция как тип
type MyFuncType func(int, int) int

// MyTypeForCallBack функция как тип
type MyTypeForCallBack func(a int, b int)

// MyCallBack Функция для колбека
func MyCallBack(a int, b int) {
	fmt.Println("Hello from MyCallBack function", a, b)
}

// FuncWithCallBack функция обратного вызова
func FuncWithCallBack(a int, b int, f MyTypeForCallBack) {
	f(a, b)
}

// Prefixer Замыкание
type Prefixer func(string) string

// DeferExample defer
func DeferExample() {
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	fmt.Println("some work.....")
}

// PanicExample panic
func PanicExample() {
	panic("panic")
}

// RecoverExample recover
func RecoverExample() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	panic("panic")
}

func RecoverExample2() {
	defer func(a int) {
		fmt.Println("the last", a)
	}(25)

	fmt.Println("some work.....")

	panic("panic")
}

func DeferPanicsTest() {
	defer func() {
		fmt.Println("defer 1")
	}()

	defer func() {
		fmt.Println("defer 2")
	}()

	defer func() {
		fmt.Println("defer 3")
	}()

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic", r)
		}
	}()

	panic("panic 1")
}
