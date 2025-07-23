package main

import (
	"fmt"

	"github.com/k0kubun/pp/v3"
)

type Message struct {
	id          int
	text        string
	address     string
	typeChannel string
}

func main() {
	sliceMessage := []Message{
		Message{
			id:          1,
			text:        "new message!",
			address:     "89193334455",
			typeChannel: "sms",
		},
		Message{
			id:          1,
			text:        "new order!",
			address:     "89193546245",
			typeChannel: "sms",
		},
		Message{
			id:          1,
			text:        "new discount!",
			address:     "89034459323",
			typeChannel: "sms",
		},
	}

	for _, value := range sliceMessage {
		pp.Println(value)
	}

	fmt.Println("=== ADDED NEW MESSAGE ===")

	sliceMessage = append(sliceMessage, Message{
		id:          4,
		text:        "new salary: 4200$",
		address:     "86064433432",
		typeChannel: "sms",
	})

	for _, value := range sliceMessage {
		pp.Println(value)
	}
	fmt.Println("Капасити (вместимость слайса):", cap(sliceMessage))
}
