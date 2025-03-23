package main

import (
	"fmt"
	"os"
)

var saidhello bool
var current_target string

func main() {
	hello()
	if saidhello {

		fmt.Println("Кого мы ищем сегодня?")
		fmt.Scanln(&current_target)

		fmt.Println("Тогда, наша цель -", current_target)
	}
	filecheck()
	ask()

}

func hello() {
	fmt.Println("Салам алейкум всем достойным!")
	saidhello = true
}

func filecheck() {
	fileinfo, err := os.Stat(current_target)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	} else {
		fmt.Println("---------------------------------------- \n Имя файла -", fileinfo.Name(), "\n Его размер -", fileinfo.Size(), "байт \n----------------------------------------")

		// ask(decision) //А потом здесь ставится новый блок if{}, в котором, в зависимости от решения, будет выполняться действие.
		return
	}
}

func ask() {
	var decision string
	fmt.Println("Вам оно надо(Д/Н)?")
	fmt.Scanln(&decision)
	switch decision {
	case "д":
		fmt.Println("Окей")

	case "н":
		os.Remove(current_target)
		fmt.Println("Файл", current_target, "подлежит удалению")

	default:
		fmt.Println("Чё?")
	}

}
