package main

import (
	"errors"
	"fmt"
)

const winePrice = 100

func main() {
	change, err := buyWine(18, 110)
	if err != nil {
		fmt.Println("Покупка неуспешна:", err.Error())
	} else {
		fmt.Printf("Вада сдача - %d. Хорошего дня!\n", change)
	}

	change, err = buyWine(17, 110)
	if err != nil {
		fmt.Println("Покупка неуспешна:", err.Error())
	} else {
		fmt.Printf("Вада сдача - %d. Хорошего дня!", change)
	}

	change, err = buyWine(33, 90)
	if err != nil {
		fmt.Println("Покупка неуспешна:", err.Error())
	} else {
		fmt.Printf("Вада сдача - %d. Хорошего дня!", change)
	}
}
func buyWine(age, moneyAmout int) (int, error)  {
	if age >= 18 && moneyAmout >= winePrice {
		return moneyAmout - winePrice, nil
	} else if age <18 {
		return moneyAmout, errors.New("Иди пей вишнёвый сок, сынулся")
	} else {
		return moneyAmout, errors.New("У вас недостаточно средств")
	}

}