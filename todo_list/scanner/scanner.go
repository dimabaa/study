package scanner

import (
	"bufio"
	"fmt"
	"os"
)

func PrintSS() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Вас приветствует ToDo List")
	fmt.Print("Введите команду: ")
	scanner.Scan()
	vvod := scanner.Text()
	fmt.Println("Ваша команда: ", vvod)
}
