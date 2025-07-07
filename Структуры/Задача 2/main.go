package main

import "fmt"

type User struct {
	name   string
	rating float64
}

func (u User) Greeting() {
	fmt.Println("")
	fmt.Println("Всем привет!")
	fmt.Println("Меня зовут ", u.name)
	fmt.Println("Мой рейтинг: ", u.rating)
	fmt.Println("")
}

func (u User) GoodBye() {
	fmt.Println("")
	fmt.Println("Всем пока, меня звали ", u.name)
	fmt.Println("Мой рейтинг был ", u.rating)
	fmt.Println("")
}

/*
func (u User) ratingUP(rating float64) {
	if u.rating+rating <= 10.0 {
		u.rating += rating
		fmt.Println("я добавил рейтинг. теперь он", u.rating)
	} else {
		fmt.Println("я не прошёл валидацию")
	}
}
*/
func ratingUP(u *User, rating float64) {
	if u.rating+rating <= 10.0 {
		u.rating += rating
		fmt.Println("я добавил рейтинг. теперь он", u.rating)
	} else {
		fmt.Println("я не прошёл валидацию")
	}
}

func main() {
	user := User{
		name:   "Дмитрий",
		rating: 5.0,
	}

	fmt.Println("Юзер рейтинг ДО ", user.rating)

	user.Greeting()
	user.GoodBye()

	ptr := &user

	ratingUP(ptr, 111)
	fmt.Println("Юзер рейтинг ПОСЛЕ", user.rating)
}
