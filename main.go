package main

import (
	"app/account"
	"app/files"
	"fmt"
)

func main() {
cycle:
	for {
		fmt.Println("Выберите что вы хотите сделать: ")
		menu()
		var chouseUser int
		fmt.Scan(&chouseUser)

		switch chouseUser {
		case 1:
			createAccount()
		case 2:
			//
		case 3:
			//
		case 4:
			break cycle
		}
	}
}

func createAccount() {

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

	file, err := account.ToBytes()
	if err != nil {
		fmt.Println("Не удалось преобразовать файл в JSON")
		return
	}
	files.WriteFile(file, "data.json")
}

func printData(promt string) string {
	fmt.Println(promt)
	var res string
	fmt.Scanln(&res)
	return res
}

func menu() {
	fmt.Println("1. Создать аккаунт")
	fmt.Println("2. Найти аккаунт")
	fmt.Println("3. Удалить аккаунт")
	fmt.Println("4. Выход")
}
