package main

import "fmt"

func main() {

	people := make(map[string]bool, 3)

	people["Иван Иванов"] = true
	people["Сергей Петров"] = false

	c, ok := people["Сергей Петров"]

	if !ok {
		fmt.Println("people NOT found in database")
		return
	}

	if ok {
		if c {
			fmt.Println("people is Criminal")
		} else {
			fmt.Println("people is NOT criminal")
		}
	}
	fmt.Println(c, ok)
}
