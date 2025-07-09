package interfaces

type IAnimal interface {
	Walking()
	Eating()
}

type IDogLike interface {
	IAnimal
	Howling()
}

type ICatLike interface {
	IAnimal
	Roaring()
}
