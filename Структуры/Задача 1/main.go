package main

import "fmt"

type User struct {
	name       string
	age        int
	phone      int
	isClose    bool
	rating     float64
	profession string
}

type Message struct {
	dateTimeCreate string
	dateTimeStop   string
	channel        int
	system         string
	address        string
	subject        string
	body           string
	status         int
}

func main() {
	message := Message{
		dateTimeCreate: "27.06.2025 16:44:23",
		dateTimeStop:   "27.06.2025 17:44:23",
		channel:        1,
		system:         "dbo",
		address:        "9197665337",
		subject:        "Платёжная операция",
		body:           "Произведена оплата в магазине ВкусВилл на сумму 1342.50 руб",
		status:         1,
	}
	user := User{
		name:       "Дмитрий",
		age:        33,
		phone:      9197665337,
		isClose:    false,
		rating:     5.0,
		profession: "golang engeneer",
	}
	fmt.Println(message)
	fmt.Println(user)

	fmt.Println("Адрес до", message.address)
	message.address = "7771234568"
	fmt.Println("Адрес после", message.address)

}
