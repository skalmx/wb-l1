package main

import "fmt"

type Human struct {
	name    string
	surname string
}

func (h *Human) SetName(name string) {
	h.name = name
}

type Action struct {
	Human
	age int
}

func (a *Action) SetAge(age int) {
	a.age = age
}

func main() {
	Human := &Human{
		name:    "Nikita",
		surname: "Surname",
	}
	fmt.Println(Human) // Nikita Surname
	Action := &Action{Human: *Human, age: 20}
	fmt.Println(Action) // Nikita Surname 20
	Action.SetAge(18)
	Action.SetName("TestName")
	fmt.Println(Action) // TestName Surname 18

}