package factory

import (
	"practice/designPatterns/factoryPatterns/animals/interfaces"
	"practice/designPatterns/factoryPatterns/animals/structs"
)

type AnimalFactoryLion struct{}

func (f *AnimalFactoryLion) CreateAnimal(name string, age, weight int) interfaces.IAnimal {
	return structs.NewLion(name, age, weight)
}

type AnimalFactoryWolf struct{}

func (f *AnimalFactoryWolf) CreateAnimal(name string, age, weight int) interfaces.IAnimal {
	return structs.NewWolf(name, age, weight)
}

type AnimalFactoryLeopard struct{}

func (f *AnimalFactoryLeopard) CreateAnimal(name string, age, weight int) interfaces.IAnimal {
	return structs.NewLeopard(name, age, weight)
}
