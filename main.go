package main

import (
	"app/account"
	"app/files"
	"fmt"
)

func main() {
	files.ReadFile()
	files.WriteFile("Одаб", "dabdab.txt")

	login := printData("Введите свой логин")

	var password string
	fmt.Println("Введите свой пароль")
	fmt.Scanln(&password)

	url := printData("Введите свой url")

	account, err := account.NewAccount(login, password, url)

	if err != nil {
		fmt.Println("Неверный формат URL или login!")
		return
	}

	account.OutputAccount()

}

func printData(promt string) string {
	fmt.Println(promt)
	var res string
	fmt.Scanln(&res)
	return res
}
