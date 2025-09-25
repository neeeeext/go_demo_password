package account

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"net/url"
)

// const
var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$")

//Struct

type Account struct {
	Login    string
	Password string
	Url      string
}

func NewAccount(login, password, urlString string) (*Account, error) {
	if login == "" {
		return nil, errors.New("INVALID_URL")

	}
	_, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, errors.New("INVALID_URL")
	}

	newAccount := &Account{
		Login:    login,
		Password: password,
		Url:      urlString,
	}
	if password == "" {
		newAccount.generatePassword()
	}
	return newAccount, nil
}

//Method

func (a *Account) OutputAccount() {
	fmt.Println(a.Login, a.Password, a.Url)
}

func (a *Account) generatePassword() {

	fmt.Println("Введите сколько знаков в пароле хотите")
	var rndNumber int
	fmt.Scan(&rndNumber)

	password := make([]rune, rndNumber)

	for i := 0; i < rndNumber; i++ {
		password[i] = letter[rand.IntN(len(letter))]

	}

	a.Password = string(password)
	fmt.Println("Ваш пароль сгенирован и записан")
}
