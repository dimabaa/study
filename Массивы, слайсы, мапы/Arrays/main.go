package main

import (
	"fmt"

	"github.com/k0kubun/pp/v3"
)

type User struct {
	Name    string
	Rating  float64
	Premium bool
}

func main() {

	userArray := [3]User{
		User{
			Name:    "Dima",
			Rating:  5.0,
			Premium: false,
		},
		User{
			Name:    "Daniil",
			Rating:  7.9,
			Premium: true,
		},
		User{
			Name:    "Arina",
			Rating:  9.0,
			Premium: true,
		},
	}

	fmt.Println(userArray)

	fmt.Println("Before:")
	fmt.Println("--------------------")

	for _, user := range userArray {
		pp.Println(user)
	}

	/*
		for i := 0; i < len(userArray); i++ {
			pp.Println(userArray[i])
		}
	*/
	fmt.Println("")

	// for i := 0; i < len(userArray); i++ {
	// 	if userArray[i].Premium {
	// 		userArray[i].Rating += 1.0
	// 	}
	// }

	for index, value := range userArray {
		if value.Premium {
			userArray[index].Rating += 1.5
		}
	}

	fmt.Println("After:")
	fmt.Println("--------------------")
	// for i := 0; i < len(userArray); i++ {
	// 	pp.Println(userArray[i])
	// }

	for index, value := range userArray {
		fmt.Println("Индекс:", index)
		pp.Println("Значение:", value)
	}

}
