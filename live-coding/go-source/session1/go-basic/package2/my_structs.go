package package2

import "fmt"

type Species interface {
	Describe() string
}

type Human struct {
	name   string
	age    int
	wealth *int
}

func (h Human) Describe() string {
	return fmt.Sprintf("I am a human named %s and I am %d years old", h.name, h.age)
}

type Animal struct {
	Name string
	Age  int
}

func (a Animal) Describe() string {
	return fmt.Sprintf("I am an animal named %s and I am %d years old", a.Name, a.Age)
}

func (h *Human) SetName(name string) {
	h.name = name
}

func (h *Human) SetAge(age int) {
	h.age = age
}

func (h *Human) SetWealth(wealth *int) {
	h.wealth = wealth
}

func (h Human) String() string {
	return fmt.Sprintf("Name: %s, Age: %d", h.name, h.age)
}

func (h Human) IncorrectSetAge(age int) {
	h.age = age
}

func (h Human) IncorrectSetName(name string) {
	h.name = name
}
