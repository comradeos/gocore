package main

import f "fmt"
import h "go003Basics/helper"

func main() {
	f.Println("Maps(Dicts):")
	h.Sep()

	// створення мапи
	var ages map[string]int
	ages = make(map[string]int)
	// створення мапи

	f.Println(ages)
	h.Sep()

	ages["Василь"] = 20
	ages["Марічка"] = 30
	ages["Петро"] = 40
	ages["Ганна"] = 50
	
	f.Println(ages)
	h.Sep()

	for key, value := range ages {
		f.Printf("%s -> %d\n", key, value)
	}

	var langs = map[int]string {
		2018: "Php",
		2019: "Python",
		2020: "C/C++",
		2021: "C#",
		2022: "Go",
	}

	f.Println(langs)
	f.Println(langs[2022])

	h.Sep()

	value, exists := langs[2090]
	f.Printf("value = %s, exists = %t\n", value, exists)

	value, exists = langs[2022]
	f.Printf("value = %s, exists = %t\n", value, exists)

	h.Sep()

	delete(langs, 2018)
	f.Println(langs)
}