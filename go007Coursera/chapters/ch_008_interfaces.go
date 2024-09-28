package chapters

import "fmt"

func Ch008Interfaces() {
	fmt.Println("============== Hello from Ch008Interfaces() ==============")

	// Интерфейс - это набор методов, которые должен реализовать тип данных.

	// Пример интерфейса Payer:

	var wallet Payer = &Wallet{Cash: 1}

	fmt.Println(wallet)

	// Пример использования интерфейса:

	err := wallet.Pay(10)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Payment success")
	}

	// Пример использования интерфейса с другим типом данных

	var iaroslav Payer = &Iaroslav{Name: "Iaroslav", Cash: 100}
	_ = iaroslav.Pay(10)

	var wayToPay Payer = wallet
	_ = wayToPay.Pay(10)

}

type Payer interface {
	Pay(int) error
}

type Wallet struct {
	Cash int
}

func (w *Wallet) Pay(amount int) error {
	if w.Cash < amount {
		return fmt.Errorf("not enough money")
	}
	w.Cash -= amount
	return nil
}

type Iaroslav struct {
	Name string
	Cash int
}

func (i *Iaroslav) Pay(amount int) error {
	if i.Cash < amount {
		return fmt.Errorf("iaroslav has not enough money")
	} else {
		fmt.Println("Iaroslav has enough money. All is good.")
	}
	return nil
}
