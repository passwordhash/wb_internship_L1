package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

func (h Human) Introduce() {
	fmt.Printf("Hello, my name is %s, I am %d years old\n", h.Name, h.Age)
}

type Action struct {
	ActionTitle string
	Human
}

func (a Action) Act() {
	fmt.Printf("%s is doing %s\n", a.Name, a.ActionTitle)
}

func main() {
	publisher := Action{
		Human:       Human{"Yaroslav", 20},
		ActionTitle: "content publishing",
	}

	publisher.Introduce()
	publisher.Act()
}
