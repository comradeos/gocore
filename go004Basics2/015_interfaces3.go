package main // Определение пакета main, где начинается выполнение программы

import (
    "fmt" // Импорт пакета fmt для вывода на консоль
)

func main() { // Главная функция, где начинается выполнение программы
    fmt.Println("Interfaces")

    var user IUser = NewIUser("John", "john@mail.com") // Создание нового пользователя с помощью функции NewIUser
    user.Info() // Вызов метода Info для пользователя

	fmt.Printf("%v, %T\n", user, user) // Вывод информации о пользователе

	var a *int
	if a == nil {
		fmt.Println("a is nil")
	} else {
		fmt.Println("a is not nil")
	}
}

type user struct { // Определение структуры user
    Name string // Поле Name структуры user
    Email string // Поле Email структуры user
}

func (u *user) Info() { // Метод Info для структуры user, который выводит информацию о пользователе
    fmt.Println(u.Name, u.Email)
}

func (u *user) ChangeName(newName string) { // Метод ChangeName для структуры user, который изменяет имя пользователя
    u.Name = newName
}

func (u *user) ChangeEmail(newEmail string) { // Метод ChangeEmail для структуры user, который изменяет email пользователя
    u.Email = newEmail
}

type IUser interface { // Определение интерфейса IUser
    Info() // Метод Info, который должен быть реализован в любом типе, который реализует интерфейс IUser
    ChangeName(newName string) // Метод ChangeName, который должен быть реализован в любом типе, который реализует интерфейс IUser
    ChangeEmail(newEmail string) // Метод ChangeEmail, который должен быть реализован в любом типе, который реализует интерфейс IUser
}

func NewIUser(name, email string) IUser { // Функция NewIUser, которая создает нового пользователя и возвращает его как тип IUser
    return &user{name, email}
}