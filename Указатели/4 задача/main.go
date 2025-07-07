package main

import "fmt"

func main() {
	var size1 float64
	fmt.Println("введи текущий размер груди: ")
	fmt.Scan(&size1)

	//size := 2.0
	fmt.Println("исходный размер ", size1)
	ptrSize := &size1
	fmt.Println("Создали указатель на размер, значение: ", *ptrSize, " адрес: ", ptrSize)
	changeGrud(ptrSize)
	fmt.Println("Ну и вызываем снова старый размер, а он уже тоже ", size1)
}

func changeGrud(a *float64) {
	if *a == 3 {
		fmt.Println("У вас отличный размер груди <", *a, "- ка> не много и не мало!")
		fmt.Println("Оставляем как есть!")
		return
	}
	if *a < 3 {
		*a = 3.5
		fmt.Println("в функции изменили размер груди, теперь он", *a)
		return
	}

	if *a > 3 {
		*a = 3
		fmt.Println("в функции изменили размер груди, теперь он", *a)
		return
	}

}
