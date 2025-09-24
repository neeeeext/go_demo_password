package main

import (
	"fmt"
)

type account struct {
	Login    string
	Password string
	Url      string
}

func main() {
	for {
		login := printData("Введите свой логин")
		password := printData("Введите свой пароль")
		url := printData("Введите свой url")

		account := account{
			Login:    login,
			Password: password,
			Url:      url,
		}
		outputAccount(&account)

	}

}

func printData(promt string) string {
	fmt.Println(promt)
	var res string
	fmt.Scan(&res)
	return res
}

func outputAccount(a *account) {
	fmt.Println(a.Login)
	fmt.Println(a.Password)
	fmt.Println(a.Url)
}
