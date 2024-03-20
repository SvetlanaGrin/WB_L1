package main

import "fmt"

type Human struct {
	name string
	age  int
}

func (h *Human) change() {
	fmt.Println("Сменить действие")
}
func NewHuman(name string, age int) *Human {
	return &Human{name: name, age: age}
}
func (h *Human) move() {
	fmt.Println("Ходить")
}

type Action struct {
	Human
}

func (a *Action) move() {
	fmt.Println("Бежать")
}
func main() {
	human := NewHuman("Svet", 20)
	human.move()
	human.change()
	action := Action{*human}
	action.move()
	action.change()
	human.move()
}
