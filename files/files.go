package files

import (
	"fmt"
	"os"
)

func WriteFile(content []byte, name string) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println("Ошибка создания файла")
	}
	defer file.Close()

	_, err = file.Write(content)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Запись успешна")
}

func ReadFile() {
	data, err := os.ReadFile("dabdab.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(data))
}
