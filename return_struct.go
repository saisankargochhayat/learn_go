package main

import "fmt"

type Person interface {
	Name() string
	Age() int
}

type Game string
type Gamer struct {
	name  string
	age   int
	games []Game
}

func NewGamer(name string, age int) *Gamer {
	return &Gamer{
		name: name,
		age:  age,
	}
}

func (g *Gamer) Name() string {
	return g.name
}

func (g *Gamer) Age() int {
	return g.age
}

func (g *Gamer) Games() []Game {
	return g.games
}

func (g *Gamer) AddGame(game Game) {
	g.games = append(g.games, game)
}

func main() {
	var person Person = NewGamer("Jane Doe", 42)
	fmt.Println(person.Name())

	var gamer *Gamer = NewGamer("John Doe", 29)
	gamer.AddGame("Metal Gear Solid 3")
	fmt.Println(gamer.Games())
}
