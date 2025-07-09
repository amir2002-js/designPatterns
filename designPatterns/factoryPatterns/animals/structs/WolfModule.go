package structs

import (
	"fmt"
	"practice/designPatterns/factoryPatterns/animals/interfaces"
)

type Wolf struct {
	Animals
}

func (w *Wolf) Walking() {
	fmt.Printf("%s is walking\n", w.name)
}

func (w *Wolf) Eating() {
	fmt.Printf("%s is eating\n", w.name)
}

func (w *Wolf) Howling() {
	fmt.Printf("%s is roaring\n", w.name)
}

func NewWolf(name string, age, weight int) interfaces.IDogLike {
	return &Wolf{Animals{name: name, animalType: "dog", age: age, weight: weight}}
}
