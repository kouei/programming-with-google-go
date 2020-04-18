package main

import (
	"fmt"
	"os"
)

// Animal interface
type Animal interface {
	Eat()
	Move()
	Speak()
}

// Cow type
type Cow struct {
	name string
}

// Eat method for cow
func (c Cow) Eat() {
	fmt.Print("grass")
}

// Move method for cow
func (c Cow) Move() {
	fmt.Print("walk")
}

// Speak method for cow
func (c Cow) Speak() {
	fmt.Print("moo")
}

// Bird type
type Bird struct {
	name string
}

// Eat method for bird
func (b Bird) Eat() {
	fmt.Print("worms")
}

// Move method for bird
func (b Bird) Move() {
	fmt.Print("fly")
}

// Speak method for bird
func (b Bird) Speak() {
	fmt.Print("peep")
}

// Snake Type
type Snake struct {
	name string
}

// Eat method for snake
func (s Snake) Eat() {
	fmt.Print("mice")
}

// Move method for snake
func (s Snake) Move() {
	fmt.Print("slither")
}

// Speak method for snake
func (s Snake) Speak() {
	fmt.Print("hsss")
}

func readString() string {
	var s string
	nRead, err := fmt.Scan(&s)
	if nRead != 1 || err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return s
}

func newAnimal(m *map[string]Animal, name string) {
	animalType := readString()

	if animalType == "cow" {
		(*m)[name] = Cow{name}
	} else if animalType == "bird" {
		(*m)[name] = Bird{name}
	} else if animalType == "snake" {
		(*m)[name] = Snake{name}
	} else {
		fmt.Println("Unknown Type")
		return
	}
	fmt.Printf("New %s added: %s\n", animalType, name)
}

func query(m *map[string]Animal, name string) {
	animal, isExist := (*m)[name]
	if !isExist {
		fmt.Printf("Unknown Name [%s]\n", name)
		return
	}

	info := readString()
	if info == "eat" {
		fmt.Printf("The food of [%s] is [", name)
		animal.Eat()
		fmt.Println("]")
	} else if info == "move" {
		fmt.Printf("The locomotion of [%s] is [", name)
		animal.Move()
		fmt.Println("]")
	} else if info == "speak" {
		fmt.Printf("The noise of [%s] is [", name)
		animal.Speak()
		fmt.Println("]")
	} else {
		fmt.Printf("Unknown Info [%s]\n", info)
		return
	}
}

func main() {

	m := map[string]Animal{}

	for {
		fmt.Print("> ")
		command := readString()
		name := readString()

		if command == "newanimal" {
			newAnimal(&m, name)
		} else if command == "query" {
			query(&m, name)
		} else {
			fmt.Println("Unknown Command")
		}
	}
}
