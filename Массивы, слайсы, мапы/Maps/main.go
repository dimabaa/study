package main

import "fmt"

func main() {
	weather := map[int]int{
		11: +3,
		12: +6,
		13: +0,
		14: -4,
		15: +1,
		30: 0,
	}
	fmt.Println(weather[12])
	c, ok := weather[30]
	fmt.Println(c, ok)

	fmt.Println(weather[30])
	fmt.Println(weather[40])
	weather[20] = -10
	fmt.Println(weather[20])
}
