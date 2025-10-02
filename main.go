package main

import (
	"app/account"
	"app/files"
	"app/output"
	"fmt"

	"github.com/fatih/color"
)

func main() {
	vault := account.NewVault(files.NewJsonDb("data.json"))
Menu:
	for {
		chouseUser := menu()

		switch chouseUser {
		case 1:
			createAccount(vault)
		case 2:
			findAccount(vault)
		case 3:
			deleteAccount(vault)
		default:
			break Menu
		}
	}
}
func findAccount(vault *account.VaultWithDb) {
	findUrl := printData("Введите url по которому хотите найти пароль")

	isTrueAccounts := vault.FindAccountsByUrl(findUrl)
	if len(isTrueAccounts) == 0 {
		output.PrintError("Нужный аккаунт не найден!")
	}
	for _, account := range isTrueAccounts {
		account.OutputAccount()
		fmt.Println("")
	}

}
func deleteAccount(vault *account.VaultWithDb) {
	findUrl := printData("Введите url по которому хотите найти пароль")
	isDeleted := vault.DeleteAccountByUrl(findUrl)
	if isDeleted {
		color.Green("Удалено")
	} else {
		output.PrintError("Не найдено")
	}
}

func createAccount(vault *account.VaultWithDb) {

	login := printData("Введите свой логин")

	var password string
	fmt.Println("Введите свой пароль")
	fmt.Scanln(&password)

	url := printData("Введите свой url")

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println("Неверный формат URL или login!")
		return
	}

	vault.AddAccount(*myAccount)

}

func printData(promt string) string {
	fmt.Println(promt)
	var res string
	fmt.Scanln(&res)
	return res
}

func menu() int {
	fmt.Println("Выберите что вы хотите сделать: ")
	fmt.Println("")
	fmt.Println("1. Создать аккаунт")
	fmt.Println("2. Найти аккаунт")
	fmt.Println("3. Удалить аккаунт")
	fmt.Println("4. Выход")

	var chouseUser int
	fmt.Scan(&chouseUser)
	return chouseUser
}
