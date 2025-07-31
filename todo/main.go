package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Вас приветствует ToDo List")
	fmt.Print("Введите команду")
	vvod := scanner.Text()
	fmt.Println(vvod)

}
