package main

import "fmt"

func main() {
	intSlice := make([]int, 0, 10)

	intSlice = append(intSlice, 1, 2, 3, 4, 5, 6)
	intSlice = append(intSlice, 1)

	fmt.Println(intSlice)
	fmt.Println("cap:", cap(intSlice))
	fmt.Println("len", len(intSlice))
}
