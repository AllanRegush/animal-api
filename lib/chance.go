package lib

import "math/rand"

type Chance struct{}

func (*Chance) Name() string {
	names := []string{"Bob", "Tommy", "Anne", "Mary"}

	name := names[rand.Intn(len(names))]

	return name
}

func (*Chance) Age() int {
	return rand.Intn(100) + 1
}

func (*Chance) Animal() string {
	animals := []string{"Tiger", "Monkey", "Cat", "Dog", "Deer", "Snake"}

	animal := animals[rand.Intn(len(animals))]
	return animal
}
