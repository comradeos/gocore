package chapters

import "fmt"

func Ch009EmptyInterface() {
	fmt.Println("============== Hello from Ch009EmptyInterface() ==============")

	// Пустой интерфейс - это интерфейс, у которого нет методов.
	var wallet1 Wallet1 = Wallet1{Cash: 100}

	fmt.Printf("1) Raw payment %#v\n", wallet1)      // Raw payment chapters.Wallet1{Cash:100}
	fmt.Printf("2) Formatted payment %v\n", wallet1) // Formatted payment {100}
	fmt.Printf("3) %s\n", wallet1)

	Buy(wallet1) // Wallet1 with cash: 100 dollars
	Buy(19321)
}

type Wallet1 struct {
	Cash int
}

func (w Wallet1) String() string {
	return "Wallet1 with cash: " + fmt.Sprint(w.Cash) + " dollars"
}

type Payer1 interface {
	Pay(int) error
}

func (w Wallet1) Pay(amount int) error {
	if w.Cash < amount {
		return fmt.Errorf("not enough money")
	}
	w.Cash -= amount
	return nil
}

func Buy(i interface{}) {
	var p Payer1
	var ok bool

	if p, ok = i.(Payer1); !ok {
		fmt.Printf("%T is not a payer\n", i)
		return
	}

	err := p.Pay(10)
	if err != nil {
		fmt.Printf("Payment failed: %v\n", err)
		return
	}

	fmt.Printf("%T is a payer\n", i)
}

// преобразование интерфейсов пример
