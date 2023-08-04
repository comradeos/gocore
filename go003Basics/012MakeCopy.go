package main

import f "fmt"
import h "go003Basics/helper"

func main() {
	f.Println("MAKE() & COPY()")
	h.Sep()

	var s = []int { 1, 2, 3, 4, 5, 6, 7 }
	f.Println(s)
	
	var s1 = s[3:6] 
	f.Println(s1)

	s1[0] = 777
	f.Println(s)



	// linking... not copy
	h.Sep()
	a := s
	a[0] = 222
	f.Println(a)
	f.Println(s)


	h.Sep()
	// copy of slice
	var cp = make([]int, len(s))
	copy(cp, s)
	cp[0] = 1
	f.Println(cp)
	f.Println(s)

	
}