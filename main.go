package main

import (
	"app/account"
	"app/files"
	"app/output"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

var menu = map[string]func(*account.VaultWithDb){
	"1": createAccount,
	"2": findAccountByUrl,
	"3": findAccountByLogin,
	"4": deleteAccount,
}

func main() {
	res := os.Getenv("VAR")
	fmt.Println(res)

	err := godotenv.Load()
	if err != nil {
		output.PrintError("Не удалось загрузить env файл")
	}

	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
	}

	vault := account.NewVault(files.NewJsonDb("data.json"))
Menu:
	for {
		variant := printData(
			"1. Создать аккаунт",
			"2. Найти аккаунт по URL",
			"3. Найти аккаунт по Login",
			"4. Удалить аккаунт",
			"5. Выход",
			"Выберите что вы хотите сделать: ",
		)

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
func findAccountByUrl(vault *account.VaultWithDb) {
	findUrl := printData("Введите url по которому хотите найти аккаунт")

	isTrueAccounts := vault.FindAccounts(findUrl, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Url, str)

	})
	findAccountOutput(&isTrueAccounts)

}

func findAccountByLogin(vault *account.VaultWithDb) {
	findLogin := printData("Введите login по которому хотите найти аккаунт")

	isTrueAccounts := vault.FindAccounts(findLogin, func(acc account.Account, str string) bool {
		return strings.Contains(acc.Login, str)

	})
	findAccountOutput(&isTrueAccounts)
}

func findAccountOutput(a *[]account.Account) {
	if len(*a) == 0 {
		output.PrintError("Нужный аккаунт не найден!")
	}
	for _, account := range *a {
		account.OutputAccount()
		fmt.Println("")
	}
}

func deleteAccount(vault *account.VaultWithDb) {
	findUrl := printData("Введите url по которому хотите УДАЛИТЬ аккаунт")
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

	url := printData([]string{"Введите свой url"})

	myAccount, err := account.NewAccount(login, password, url)
	if err != nil {
		fmt.Println("Неверный формат URL или login!")
		return
	}

	vault.AddAccount(*myAccount)

}

func printData(promt ...any) string {
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
