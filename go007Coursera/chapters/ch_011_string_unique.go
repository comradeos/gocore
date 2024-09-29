package chapters

import (
	"bufio"
	"fmt"
	"os"
)

func Ch011StringUnique() {
	fmt.Println("============== Hello from Ch01StringUnique() ==============")

	var seen = make(map[string]bool)

	fmt.Printf("Enter text: ")

	in := bufio.NewScanner(os.Stdin)

	for in.Scan() {

		text := in.Text()

		if _, ok := seen[text]; ok {
			fmt.Println("Already seen")
			continue
		}

		fmt.Printf("Text: %s\n", text)
		seen[text] = true
	}

}
