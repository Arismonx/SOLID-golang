package main

import (
	"fmt"
	"time"
)

type (
	Animal interface {
		Speak() string
	}

	Dog  struct{}
	Cat  struct{}
	Bird struct{}
)

func (b Bird) Speak() string {
	return "Tweet!"
}

func (d Dog) Speak() string {
	return "Woof!"
}
func (c Cat) Speak() string {
	return "Meow!"
}

func MakeAnimalSpeak(a Animal) {
	fmt.Println(a.Speak())
}

func main() {
	start := time.Now()
	dog := Dog{}
	cat := Cat{}
	bird := Bird{}
	// fmt.Println(dog.Speak())
	// fmt.Println(cat.Speak())
	MakeAnimalSpeak(dog)
	MakeAnimalSpeak(cat)
	MakeAnimalSpeak(bird)

	fmt.Println("TIME:", time.Since(start))
}
