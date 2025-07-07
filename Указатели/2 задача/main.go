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

	var q1 *int
	fmt.Println(q1)
	var q2 *string
	fmt.Println(q2)
	var q3 *bool
	fmt.Println(q3)
	var q4 *float64
	fmt.Println(q4)

	provInt(ptr1)
	provInt(q1)

	provStr(ptr2)
	provStr(q2)

	provBool(ptr3)
	provBool(q3)

	provFloat(ptr4)
	provFloat(q4)

}

func provInt(st *int) {
	if st != nil {
		fmt.Println("стринг равен ", *st, "его адрес ", st)
	} else {
		fmt.Println("стринг нул")
	}
}

func provStr(st *string) {
	if st != nil {
		fmt.Println("стринг равен ", *st, "его адрес ", st)
	} else {
		fmt.Println("стринг нул")
	}
}

func provBool(st *bool) {
	if st != nil {
		fmt.Println("стринг равен ", *st, "его адрес ", st)
	} else {
		fmt.Println("стринг нул")
	}
}

func provFloat(st *float64) {
	if st != nil {
		fmt.Println("стринг равен ", *st, "его адрес ", st)
	} else {
		fmt.Println("стринг нул")
	}
}
