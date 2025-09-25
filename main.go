package main

import (
	"fmt"
)

func main() {

	for {

		login := printData("Введите свой логин")

		var password string
		fmt.Println("Введите свой пароль")
		fmt.Scanln(&password)

		url := printData("Введите свой url")

		account, err := newAccount(login, password, url)

		if err != nil {
			fmt.Println("Неверный формат URL или login!")
			return
		}

		account.outputAccount()

	}
}
