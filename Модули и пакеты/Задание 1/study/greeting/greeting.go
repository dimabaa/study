package greeting

import "fmt"

func SayHello() {
	fmt.Println("Hello in GREETING")
	j := GiveMeInt()
	fmt.Println(j)
}
