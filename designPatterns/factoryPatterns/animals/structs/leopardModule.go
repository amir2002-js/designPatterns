package structs

import (
	"fmt"
	"practice/designPatterns/factoryPatterns/animals/interfaces"
)

type Leopard struct {
	Animals
}

func (l *Leopard) Walking() {
	fmt.Printf("%s is walking\n", l.name)
}

func (l *Leopard) Eating() {
	fmt.Printf("%s is eating\n", l.name)
}

func (l *Leopard) Roaring() {
	fmt.Printf("%s is roaring\n", l.name)
}

func NewLeopard(name string, age, weight int) interfaces.ICatLike {
	return &Leopard{Animals{name: name, animalType: "Cat", age: age, weight: weight}}
}
