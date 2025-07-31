package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/k0kubun/pp"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter command: ")

		if ok := scanner.Scan(); !ok {
			fmt.Println("Error entered")
			return
		}
		text := scanner.Text()

		fields := strings.Fields(text)

		if len(fields) == 0 {
			fmt.Println("Вы ничего не ввели")
			return
		}

		pp.Println(fields)

		cmd := fields[0]

		if cmd == "добавить" {
			str := ""
			for i := 1; i < len(fields); i++ {
				str += fields[i]
				if i < len(fields)-1 {
					str += "\n"
				}
			}
			fmt.Println("Вы добавили:\n", str)
		} else if cmd == "удалить" {
			str := ""
			for i := 1; i < len(fields); i++ {
				str += fields[i]
				if i < len(fields)-1 {
					str += "\n"
				}
			}
			fmt.Println("Вы удалили:\n", str)
		} else if cmd == "help" {
			fmt.Println("----------------")
			fmt.Println("Это справка")
			fmt.Println("Команды:")
			fmt.Println("<добавить> + {перечислите продукты}")
			fmt.Println("<удалить> + {перечислите продукты}")
			fmt.Println("----------------")
		} else {
			fmt.Println("Вы ввели неизвестную команду")

		}

		if cmd == "выйти" {
			return
		}

	}

}
