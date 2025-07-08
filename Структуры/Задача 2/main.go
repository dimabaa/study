package main

import "fmt"

type User struct {
	name        string
	age         int
	phoneNumber string
	isClose     bool
	rating      float64
}

func NewUser(
	name string,
	age int,
	phoneNumber string,
	isClose bool,
	rating float64,
) User {
	if name == "" {
		return User{}
	}
	if age < 0 || age > 99 {
		return User{}
	}
	if phoneNumber == "" {
		return User{}
	}
	if rating < 0.0 || rating > 10.0 {
		return User{}
	}
	return User{
		name:        name,
		age:         age,
		phoneNumber: phoneNumber,
		isClose:     isClose,
		rating:      rating,
	}
}

func (u *User) changeName(newName string) {
	if newName != "" {
		u.name = newName
	} else {
		fmt.Println("имя не может быть пустым")
	}
}

func (u *User) changeAge(newAge int) {
	if newAge > 0 && newAge < 99 {
		u.age = newAge
	} else {
		fmt.Println("возраст таким быть не может!")
	}
}

func (u *User) changePhoneNumber(newPN string) {
	if newPN != "" {
		u.phoneNumber = newPN
	} else {
		fmt.Println("Номер телефона не может быть пустым")
	}
}

func (u *User) closeAccount() {
	u.isClose = true
}

func (u *User) openAccount() {
	u.isClose = false
}

func (u *User) ratingUP(rating float64) {
	if u.rating+rating <= 10.0 {
		u.rating += rating
	} else {
		fmt.Println("Рейтинг не может быть больше 10")
	}
}

func (u *User) ratingDOWN(rating float64) {
	if u.rating-rating > 0.0 {
		u.rating -= rating
	} else {
		fmt.Println("Рейтинг не может быть больше 10")
	}
}

func main() {
	user := NewUser(
		"Дмитрий",
		33,
		"79193345566",
		false,
		7.0,
	)

	fmt.Println("Рейтинг до:", user.rating)

	user.ratingDOWN(2.5)
	user.changeName("Гоблин")
	user.changeAge(33)
	user.changePhoneNumber("89197665388")
	user.closeAccount()

	fmt.Println("Рейтинг после:", user.rating)

	fmt.Println("\nВообще вся структура с данными\n", user)
}
