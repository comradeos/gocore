package chapters

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func Flags() {
	var n = flag.Bool("n", false, "пропуск символа новой строки")
	var sep = flag.String("s", " ", "разделитель")

	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}

func Types() {
	type Celsius float64
	type Fahrenheit float64

	var c Celsius
	var f Fahrenheit

	fmt.Println(c == 0)
	fmt.Println(f >= 0)

	log.Println(c)
}

func Task2p1() {
	var elements []string

	if len(os.Args) < 2 {
		fmt.Print("feet: ")

		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		text := input.Text()

		elements = strings.Split(text, " ")
	} else {
		elements = os.Args[1:]
	}

	var numbers []float64

	for _, element := range elements {
		number, err := strconv.ParseFloat(element, 64)
		if err == nil {
			numbers = append(numbers, number)
		}
	}

	fmt.Print("meters: ")

	for _, number := range numbers {
		fmt.Printf("%.2ff=%.2fm, ", number, number*0.3048)
	}

	fmt.Println()
}

func InitTest() {
	fmt.Println("InitTest")
}

func BitOperations() {
	x := 5
	fmt.Println(x)
	x = x << 1
	fmt.Println(x)
	x = x >> 2
	fmt.Println(x)
}

func Scopes() {
	x := 1
	fmt.Println(x)

	for i := range 10 {
		fmt.Println(x)
		fmt.Println(i)
	}

	for i := 10; i < 100; i += 10 {
		fmt.Println(i)
	}
}

func Files() {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}

	exPath := filepath.Dir(ex)
	err = os.Chdir(exPath)

	if err != nil {
		panic(err)
	}

	fmt.Println(os.Getwd())
	file, err := os.Open("files/FileA.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer func(file *os.File) {
		var err = file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

}
