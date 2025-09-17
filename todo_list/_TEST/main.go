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

func AddTask(subject string, text string) TaskStruct {
	return TaskStruct{
		Subject:    subject,
		Text:       text,
		TimeCreate: time.Now(),
		Completed:  false,
	}

}

func RemoveTask(SliceTask []TaskStruct, Subject string) {
	index := -1
	for i, task := range SliceTask {
		if task.Subject == Subject {
			index = i
			break
		}
	}
	if index != -1 {
		fmt.Println("К удалению задача с id: <", index, "> и темой:", Subject)
		pp.Println("INFO: Функция удаления ещё не реализована")
		lenSliceAfter := len(SliceTask)
		aaa := append(SliceTask[:index], SliceTask[index+1:]...)
		fmt.Println(aaa)
		fmt.Println("Длина слайса до удаления:", lenSliceAfter, "длина слайса после удаления:", len(aaa))
	} else {
		pp.Println("Задача не найдена!")
	}
}

func main() {
	fmt.Println("")
	pp.Println("Вас приветствует ToDo List")
	SliceTask := make([]TaskStruct, 0, 2)
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("")

		fmt.Println("Введите одну из команд:")
		fmt.Println("			1 = добавить")
		fmt.Println("			2 = удалить")
		fmt.Println("			3 = все задачи")
		fmt.Println("			0 = выход")
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
			new := AddTask(vvodSubject, vvodText)
			SliceTask = append(SliceTask, new)
			pp.Println("Добавлена задача: ", new)
		} else if vvod == "2" { // удалить
			fmt.Print("Введите название задачи, которую требуется удалить: ")
			scanner.Scan()
			vvodSubject := scanner.Text()
			RemoveTask(SliceTask, vvodSubject)
		} else if vvod == "3" { // задачи
			fmt.Print("Всего задач: ", len(SliceTask), " | Капасити слайса: ", cap(SliceTask))
			fmt.Println("")
			if len(SliceTask) == 0 {
				pp.Println("INFO: Ещё нет ни одной задачи!")
			} else {
				pp.Println(SliceTask)
			}

		} else if vvod == "0" { // выход
			pp.Println("До скорой встречи!")
			return
		} else {
			fmt.Println("INFO: Такой команды не существует")
		}
	}
}
