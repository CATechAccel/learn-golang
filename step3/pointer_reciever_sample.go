package main

import "fmt"

type userInterface interface {
	changeName(name string)
	changeBirth(birth string)
	changeAge(age int)
}

type user struct {
	name  string
	birth string
	age   int
}

func new() userInterface {
	return &user{
		name:  "sanae",
		birth: "5月29日",
		age:   20,
	}
}

func (u *user) changeName(name string) {
	u.name = name
}

func (u *user) changeBirth(birth string) {
	u.birth = birth
}

func (u *user) changeAge(age int) {
	u.age = age
}

func main() {

	userData := new()
	userData2 := new()

	userData.changeName("kaneda")
	userData.changeBirth("2000年")
	userData.changeAge(10)

	userData2.changeName("aaa")
	userData2.changeBirth("bbb")
	userData2.changeAge(25)

	fmt.Println(userData)
	fmt.Println(userData2)
}
