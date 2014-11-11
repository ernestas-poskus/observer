package main

import (
	"fmt"

	"github.com/ernestas-poskus/observer"
)

// Position -
type Position struct {
	x, y, z int
}

func (p *Position) String() string {
	return fmt.Sprintf("x: %d, y: %d, z: %d", p.x, p.y, p.z)
}

// NewPosition -
func NewPosition(x, y, z int) *Position {
	return &Position{x, y, z}
}

// Ball - game struct
type Ball struct {
	observer.Observable
	position *Position
}

// GetPosition - returns ball position
func (b *Ball) GetPosition() *Position {
	return b.position
}

// SetPosition - sets new ball position and notifies players
func (b *Ball) SetPosition(p *Position) {
	b.position = p
	b.Notify(p)
	fmt.Println("---------------")
}

// Player -
type Player struct {
	name         string
	lastPosition *Position
	ball         *Ball
}

// NewPlayer -
func NewPlayer(name string, b *Ball) *Player {
	p := new(Player)
	p.name = name
	p.ball = b
	return p
}

// Update -
func (p *Player) Update(values ...interface{}) {
	p.lastPosition = p.ball.GetPosition()
	fmt.Printf("%s moves with ball to %d\n", p.name, p.lastPosition)
}

func main() {

	ball := &Ball{}
	player1 := NewPlayer("Pele", ball)
	player2 := NewPlayer("Ronaldo", ball)
	player3 := NewPlayer("Zidan", ball)
	ball.Attach(player1)
	ball.Attach(player2)
	ball.Attach(player3)

	position := NewPosition(3, 22, 1)
	ball.SetPosition(position)
	ball.Detach(player3)
	ball.SetPosition(NewPosition(2281, 2911, 1))
	ball.Detach(player2)
	ball.SetPosition(NewPosition(11, 23, 1))

}
