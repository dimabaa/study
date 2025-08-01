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

	SliceTask := make([]int, 0, 10)

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
		new := Task(vvodSubject, vvodText)
		pp.Println(new)
		SliceTask = append(SliceTask, new)
	}
	fmt.Println("Ваша команда:", vvod)

}
