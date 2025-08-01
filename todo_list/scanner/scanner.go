package scanner

import (
	"bufio"
	"fmt"
	"os"
	"todo_list/task"
)

func PrintSS() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Вас приветствует ToDo List")
	fmt.Println("Введите одну из команд:")
	fmt.Println("			добавить")
	fmt.Println("			удалить")
	fmt.Println("			лист")
	fmt.Print("Ждём ввод: ")

	scanner.Scan()
	vvod := scanner.Text()
	if vvod == "добавить" {
		fmt.Print("Введите тему задачи: ")
		scanner.Scan()
		vvodSubject := scanner.Text()
		fmt.Print("Введите текст: ")
		scanner.Scan()
		vvodText := scanner.Text()
		task.Task(vvodSubject, vvodText)
	}
	fmt.Println("Ваша команда:", vvod)
}
