package main

import (
	"fmt"
	"math/rand/v2"
)

var letter = []rune("abcdefghijklmnopqrstuvwxyz")

type account struct {
	Login    string
	Password string
	Url      string
}

func main() {
	fmt.Println(len(letter))

	str := []rune("Привет!)")
	for _, ch := range string(str) {
		fmt.Println(ch, string(ch))
	}
	fmt.Println("Введите сколько знаков в пароле хотите")
	var rndNumber int
	fmt.Scan(&rndNumber)
	rndPassword := generatePassword(rndNumber)
	fmt.Println(rndPassword)

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

func generatePassword(n int) string {
	password := make([]rune, n)

	for i := 0; i < n; i++ {
		password[i] = letter[rand.IntN(len(letter))]

	}
	return string(password)
}
