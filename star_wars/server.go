//go:generate mockgen -package=$GOPACKAGE -source=$GOFILE -destination=server_mock.go
package star_wars

import "math"

type Coords struct {
	x int
	y int
}

func NewCoords(x int, y int) *Coords {
	return &Coords{
		x: x,
		y: y,
	}
}

func PlusCoords(c1 *Coords, c2 *Coords) *Coords {
	return NewCoords(
		c1.x+c2.x,
		c1.y+c2.y,
	)
}

func RotateCoords(c *Coords, angle int) *Coords {
	angleRad := float64(angle) * math.Pi / 180
	x := float64(c.x)*math.Cos(angleRad) - float64(c.y)*math.Sin(angleRad)
	y := float64(c.y)*math.Cos(angleRad) - float64(c.x)*math.Sin(angleRad)
	return NewCoords(
		int(x),
		int(y),
	)
}

type IMovable interface {
	GetPosition() *Coords
	SetPosition(position *Coords) Coords
	GetVelocity() *Coords
}

type Move struct {
	movable IMovable
}

func NewMove(movable IMovable) *Move {
	return &Move{
		movable: movable,
	}
}

func (m *Move) Execute() {
	m.movable.SetPosition(
		PlusCoords(m.movable.GetPosition(), m.movable.GetVelocity()),
	)
}

type IRotatable interface {
	GetPosition() *Coords
	GetAngle() int
	SetVelocity(velocity *Coords)
}

type Rotate struct {
	rotatable IRotatable
}

func NewRotate(rotatable IRotatable) *Rotate {
	return &Rotate{
		rotatable: rotatable,
	}
}

func (m *Rotate) Execute() {
	m.rotatable.SetVelocity(
		RotateCoords(m.rotatable.GetPosition(), m.rotatable.GetAngle()),
	)
}
