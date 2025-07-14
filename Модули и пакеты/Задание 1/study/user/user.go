package user

import "fmt"

type User struct {
	name string
	age  int
	prof string
}

func NewUser(name string, age int, prof string) User {
	if name == "" {
		fmt.Println("name не может быть пустым")
		return User{}
	}
	if age < 0 || age > 150 {
		fmt.Println("возраст не может быть:", age)
		return User{}
	}
	return User{
		name: name,
		age:  age,
		prof: prof,
	}
}

func (u *User) NewName(NewName string) {
	if NewName != "" && len(NewName) >= 2 {
		u.name = NewName
	} else {
		fmt.Println("Имя не изменено, некорректно")
	}
}

func (u *User) NewAge(NewAge int) {
	if NewAge > 0 && NewAge < 150 {
		u.age = NewAge
	}
}

func (u *User) NewProf(NewProf string) {
	if len(NewProf) > 2 && len(NewProf) < 50 {
		u.prof = NewProf
	}
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) GetAge() int {
	return u.age
}

func helper() {

}
