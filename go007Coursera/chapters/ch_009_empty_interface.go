package chapters

import "fmt"

func Ch009EmptyInterface() {
	fmt.Println("============== Hello from Ch009EmptyInterface() ==============")

	// Пустой интерфейс - это интерфейс, у которого нет методов.
	var myWallet MyWallet = MyWallet{Cash: 100}

	fmt.Printf("Raw payment %#v\n", myWallet)      // Raw payment chapters.MyWallet{Cash:100}
	fmt.Printf("Formatted payment %v\n", myWallet) // Formatted payment {100}

}

type MyWallet struct {
	Cash int
}
