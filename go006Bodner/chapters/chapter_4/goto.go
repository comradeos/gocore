package chapter_4

import "fmt"

func Goto() {
	if true {
		goto done
	}

	fmt.Println("doing...")

done: // метка
	fmt.Println("it's done")

}
