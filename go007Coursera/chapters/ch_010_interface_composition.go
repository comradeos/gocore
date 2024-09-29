package chapters

import "fmt"

func Ch010InterfaceComposition() {
	fmt.Println("============== Hello from Ch010InterfaceComposition() ==============")

	// Интерфейс можно встроить в другой интерфейс.

	// Пример интерфейса NFCPhone:
	var cellPhone Ch10NFCPhone = &Ch10CellPhone{Number: "123-456"}

	RingAndPay(cellPhone)
}

type Ch10Payer interface {
	Pay(int) error
}

type Ch10Ringer interface {
	Ring(string) error
}

type Ch10NFCPhone interface {
	Ch10Payer
	Ch10Ringer
}

type Ch10CellPhone struct {
	Number string
}

func (c *Ch10CellPhone) Pay(amount int) error {
	fmt.Println("Cell phone payment")
	return nil
}

func (c *Ch10CellPhone) Ring(number string) error {
	fmt.Println("Cell phone is ringing")
	return nil
}

func RingAndPay(phone Ch10NFCPhone) {
	_ = phone.Pay(100)
	_ = phone.Ring("123-456")
}
