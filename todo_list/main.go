package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/k0kubun/pp"
)

type TaskStruct struct {
	Subject    string
	Text       string
	TimeCreate time.Time
	Completed  bool
}

func Task(subject string, text string) TaskStruct {
	return TaskStruct{
		Subject:    subject,
		Text:       text,
		TimeCreate: time.Now(),
		Completed:  false,
	}

}
func main() {
	fmt.Println("Вас приветствует ToDo List")
	for {
		SliceTask := make([]TaskStruct, 0, 10)

		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("")

		fmt.Println("Введите одну из команд:")
		fmt.Println("			1 = добавить")
		fmt.Println("			2 = удалить")
		fmt.Println("			3 = все задачи")
		fmt.Println("			4 = выход")
		fmt.Print("Что выполняем: ")

		scanner.Scan()
		vvod := scanner.Text()
		// добавить
		if vvod == "1" {
			fmt.Print("Введите тему задачи: ")
			scanner.Scan()
			vvodSubject := scanner.Text()
			fmt.Print("Введите текст: ")
			scanner.Scan()
			vvodText := scanner.Text()
			new := Task(vvodSubject, vvodText)
			pp.Println(new)
			SliceTask = append(SliceTask, new)
		}
		// удалить
		if vvod == "2" {
			fmt.Println("Функция удаления ещё не реализована")

		}
		// задачи
		if vvod == "3" {
			pp.Println(SliceTask)

		}
		// выход
		if vvod == "4" {
			fmt.Println("До скорой встречи!")
			return
		} else {
			fmt.Println("Такой команды не существует")
		}
	}
}
