package main

import "fmt"

func main() {
	a := 11
	b := "hello"
	c := true
	d := 10.00
	fmt.Println("a", a)
	fmt.Println("b", b)
	fmt.Println("c", c)
	fmt.Println("d", d)
	ptr1 := &a
	ptr2 := &b
	ptr3 := &c
	ptr4 := &d
	fmt.Println("ptr1", ptr1)
	fmt.Println("ptr2", ptr2)
	fmt.Println("ptr3", ptr3)
	fmt.Println("ptr4", ptr4)

	fmt.Println("ptr1 znach", *ptr1)
	fmt.Println("ptr2 znach", *ptr2)
	fmt.Println("ptr3 znach", *ptr3)
	fmt.Println("ptr4 znach", *ptr4)

	changeVar1(ptr1)
	fmt.Println(*ptr1)
	changeVar2(ptr2)
	fmt.Println(*ptr2)
	changeVar3(ptr3)
	fmt.Println(*ptr3)
	changeVar4(ptr4)
	fmt.Println(*ptr4)

}

func changeVar1(q *int) {
	*q = 22
}

func changeVar2(q *string) {
	*q = "nu i dela"
}

func changeVar3(q *bool) {
	*q = false
}

func changeVar4(q *float64) {
	*q = 22.01
}
