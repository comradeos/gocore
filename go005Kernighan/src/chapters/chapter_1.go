package chapters

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"reflect"
	"strings"
	"time"
)

func Vars() {
	a := 10
	fmt.Println(a)

	var b int
	b = 20
	fmt.Println(b)

	var c = 30
	fmt.Println(c)

	var d, e int = 40, 50
	fmt.Println(d, e)

	var s string
	fmt.Println(s)
	fmt.Println(reflect.TypeOf(s))
}

func Joins() {
	var words [3]string
	words[0] = "Hello"
	words[1] = "World"
	words[2] = "!"

	var result string
	result = strings.Join(words[:], " ")
	fmt.Println(result)

	var words2 = []string{"Hello", "World", "!"}
	result = strings.Join(words2, " ")
	fmt.Println(result)

	var words3 []string
	words3 = append(words3, "Hello")
	words3 = append(words3, "World")
	words3 = append(words3, "!")
	result = strings.Join(words3, " ")
	fmt.Println(result)
}

func Maps() {
	var map1 map[string]string
	map1 = make(map[string]string)
	map1["name"] = "John"
	map1["surname"] = "Doe"
	fmt.Println(map1)
	fmt.Println(map1["name"])
}

func Dup1() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		counts[input.Text()]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d \t %s\n", n, line)
		}
	}
}

func Specifications() {
	fmt.Printf("%d\n", 10)
	fmt.Printf("%f\n", 10.0)
	fmt.Printf("%s\n", "Hello")
	fmt.Printf("%t\n", true)
}

func ScanText() {
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	text := input.Text()
	fmt.Printf("text: %s\n", text)
}

func Lissajous() {
	palette := []color.Color{color.White, color.Black}

	const (
		blackIndex = 1
		cycles     = 5
		res        = 0.001
		size       = 100
		nframes    = 64
		delay      = 8
	)

	src := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(src)

	freq := rnd.Float64() * 3.0

	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0

	for i := 0; i < nframes; i++ {
		img := image.NewPaletted(image.Rect(0, 0, 2*size+1, 2*size+1), palette)

		for t := 0.0; t < cycles*2.0; t += res {
			x := math.Sin(t * 2.0 * math.Pi)
			y := math.Sin(t*2.0*math.Pi*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}

	// Создаем файл
	f, err := os.Create("output.gif")
	if err != nil {
		fmt.Println(err)
		return
	}
	// Закрываем файл после использования
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {

		}
	}(f)

	// Записываем GIF в файл
	err = gif.EncodeAll(f, &anim)
	if err != nil {
		return
	}
}

func RandomBetween(min, max int) int {
	src := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(src)
	res := rnd.Intn(max-min) + min
	fmt.Println(res)
	return res
}

var counter int

func WebServer() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/test", test)
	err := http.ListenAndServe("localhost:8000", nil)
	if err != nil {
		return
	}
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	counter++
	_, err := fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	if err != nil {
		return
	}
}

func test(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "test %d\n", counter)
	if err != nil {
		return
	}
}
