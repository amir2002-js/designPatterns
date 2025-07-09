package structs

import (
	"fmt"
	"practice/designPatterns/factoryPatterns/animals/interfaces"
)

type Lion struct {
	Animals
}

func (l *Lion) Walking() {
	fmt.Printf("%s is walking\n", l.name)
}

func (l *Lion) Eating() {
	fmt.Printf("%s is eating\n", l.name)
}

func (l *Lion) Roaring() {
	fmt.Printf("%s is roaring\n", l.name)
}

func NewLion(name string, age, weight int) interfaces.ICatLike {
	return &Lion{Animals{name: name, animalType: "Cat", age: age, weight: weight}}
}
