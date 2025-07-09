package factory

import "practice/designPatterns/factoryPatterns/animals/interfaces"

type AnimalsFactory interface {
	CreateAnimal(name string, age, weight int) interfaces.IAnimal
}
