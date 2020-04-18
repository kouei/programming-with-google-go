package main

import (
	"fmt"
	"os"
)

// Animal describes an animal
type Animal struct {
	food, locomotion, noise string
}

// Eat returns what an animal eats
func (a Animal) Eat() string {
	return a.food
}

// Move returns how an animal moves
func (a Animal) Move() string {
	return a.locomotion
}

// Speak returns how an animal speaks
func (a Animal) Speak() string {
	return a.noise
}

func main() {

	mAnimal := map[string]Animal{
		"cow":   {"grass", "walk", "moo"},
		"bird":  {"worms", "fly", "peep"},
		"snake": {"mice", "slither", "hsss"},
	}

	mFunc := map[string]func(Animal) string{
		"eat":   Animal.Eat,
		"move":  Animal.Move,
		"speak": Animal.Speak,
	}

	mFormat := map[string]string{
		"eat":   "The food of [%s] is [%s]\n",
		"move":  "The locomotion of [%s] is [%s]\n",
		"speak": "The noise of [%s] is [%s]\n",
	}
	for {
		fmt.Print("> ")
		var name, info string
		nRead, err := fmt.Scan(&name, &info)
		if nRead != 2 || err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		animal, isExist := mAnimal[name]
		if !isExist {
			fmt.Println("Unknown Animal")
			continue
		}

		fn, isExist := mFunc[info]
		if !isExist {
			fmt.Println("Unknown Information")
			continue
		}

		format := mFormat[info]

		fmt.Printf(format, name, fn(animal))
	}
}
