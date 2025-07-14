package main

import (
	"fmt"
	"study/user"

	"github.com/k0kubun/pp/v3"
)

func main() {
	u := user.NewUser(
		"Dima",
		32,
		"Support Engeneer",
	)
	u2 := user.NewUser(
		"Petr",
		24,
		"Senior Java Developer",
	)
	fmt.Println(u, u2)
	pp.Println(u, u2)
	u.NewName("Dmitry")
	u.NewAge(33)
	u.NewProf("Golang")
	fmt.Println(u)
	fmt.Println(u.GetName())
	fmt.Println(u.GetAge())
	fmt.Println(u.GetName())
}
