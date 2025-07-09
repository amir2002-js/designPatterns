package main

import "practice/designPatterns/factoryPatterns/animals/factory"

func main() {
	lionFactory := factory.AnimalFactoryLion{}
	wolfFactory := factory.AnimalFactoryWolf{}

	lion := lionFactory.CreateAnimal("king lion", 12, 200)
	lion2 := lionFactory.CreateAnimal("horten lion", 10, 250)
	lion.Eating()
	lion2.Eating()

	wolf := wolfFactory.CreateAnimal("alpha wolf", 10, 80)
	wolf2 := wolfFactory.CreateAnimal("omega wolf", 10, 80)
	wolf2.Eating()
	wolf.Walking()
	wolf.Howling()

}
