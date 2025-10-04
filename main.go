package main

import (
	"app/account"
	"app/files"
	"app/output"
	"fmt"
	"strings"

	"github.com/fatih/color"
)

var menu = map[string]func(*account.VaultWithDb){
	"1": createAccount,
	"2": findAccount,
	"3": deleteAccount,
}

func main() {
	vault := account.NewVault(files.NewJsonDb("data.json"))
Menu:
	for {
		variant := printData([]string{
			"1. Создать аккаунт",
			"2. Найти аккаунт",
			"3. Удалить аккаунт",
			"4. Выход",
			"Выберите что вы хотите сделать: ",
		})

		menuFunc := menu[variant]
		if menuFunc == nil {
			break Menu
		}
		menuFunc(vault)

		// switch variant {
		// case "1":
		// 	createAccount(vault)
		// case "2":
		// 	findAccount(vault)
		// case "3":
		// 	deleteAccount(vault)
		// default:
		// 	break Menu
		// }
	}
}
func findAccount(vault *account.VaultWithDb) {
	findUrl := printData([]string{"Введите url по которому хотите найти пароль"})

	isTrueAccounts := vault.FindAccounts(findUrl, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Url, str)

	})
	if len(isTrueAccounts) == 0 {
		output.PrintError("Нужный аккаунт не найден!")
	}
	for _, account := range isTrueAccounts {
		account.OutputAccount()
		fmt.Println("")
	}

}

func deleteAccount(vault *account.VaultWithDb) {
	findUrl := printData([]string{"Введите url по которому хотите найти пароль"})
	isDeleted := vault.DeleteAccountByUrl(findUrl)
	if isDeleted {
		color.Green("Удалено")
	} else {
		output.PrintError("Не найдено")
	}
}

func createAccount(vault *account.VaultWithDb) {

	login := printData([]string{"Введите свой логин"})

	var password string
	fmt.Println("Введите свой пароль")
	fmt.Scanln(&password)

	url := printData([]string{"Введите свой url"})

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println("Неверный формат URL или login!")
		return
	}

	vault.AddAccount(*myAccount)

}

func printData[T any](promt []T) string {
	for i, line := range promt {
		if i == len(promt)-1 {
			fmt.Printf("%v :", line)

		} else {
			fmt.Println(line)
		}

	}
	var res string
	fmt.Scanln(&res)
	return res
}
