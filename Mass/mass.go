package main

import "fmt"

type contact struct {
	firstName string
	lastName string
	phoneNumber string
	email string
	address string
	dateofBirth string
}

func (c contact) printInfo() {
	fmt.Printf("Имя: %s\nФамилия: %s\nНомер Телефона: %s\nE-mail: %s\nАдрес: %s\nДата рождения: %s\n",
		c.firstName, c.lastName, c.phoneNumber, c.email, c.address, c.dateofBirth)
}

func (c *contact) setName(name string) {
	c.firstName = name
}

func main()  {
	c1 := contact{
		firstName: "Вася",
		lastName: "Пупкин",
		phoneNumber: "9876355112",
		email: "pulybo@mail.ru",
		address: "ул. Тутуева 15, 25",
		dateofBirth: "09.10.2001",
	}

	c1.setName("Петя")

	c1.printInfo()
}