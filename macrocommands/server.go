//go:generate mockgen -package=$GOPACKAGE -source=$GOFILE -destination=server_mock.go
package macrocommands

import (
	"errors"
)

var NoFuelError = errors.New("no fuel")

type ICommand interface {
	Execute() error
}

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

type IMovable interface {
	GetPosition() *Coords
	SetPosition(position *Coords) Coords
	GetVelocity() *Coords
}

type MoveCommand struct {
	movable IMovable
}

func NewMoveCommand(movable IMovable) *MoveCommand {
	return &MoveCommand{
		movable: movable,
	}
}

func (m *MoveCommand) Execute() error {
	m.movable.SetPosition(
		PlusCoords(m.movable.GetPosition(), m.movable.GetVelocity()),
	)
	return nil
}

type IFuelBurnable interface {
	GetLevel() int
	SetLevel(level int)
}

type CheckFuelCommand struct {
	fuelBurnable IFuelBurnable
}

func NewCheckFuelCommand(fuelBurnable IFuelBurnable) *CheckFuelCommand {
	return &CheckFuelCommand{
		fuelBurnable: fuelBurnable,
	}
}

func (m *CheckFuelCommand) Execute() error {
	level := m.fuelBurnable.GetLevel()
	if level <= 0 {
		return NoFuelError
	}

	return nil
}

type BurnFuelCommand struct {
	fuelBurnable IFuelBurnable
}

func NewBurnFuelCommand(fuelBurnable IFuelBurnable) *BurnFuelCommand {
	return &BurnFuelCommand{
		fuelBurnable: fuelBurnable,
	}
}

func (m *BurnFuelCommand) Execute() error {
	level := m.fuelBurnable.GetLevel()
	m.fuelBurnable.SetLevel(level - 1)
	return nil
}

type MacroCommand struct {
	commands []ICommand
}

func NewMacroCommand(commands []ICommand) *MacroCommand {
	return &MacroCommand{
		commands: commands,
	}
}

func (m *MacroCommand) Execute() error {
	for _, c := range m.commands {
		err := c.Execute()
		if err != nil {
			return err
		}
	}

	return nil
}

type RepeaterCommand struct {
	command ICommand
}

func NewRepeaterCommand(command ICommand) *RepeaterCommand {
	return &RepeaterCommand{
		command: command,
	}
}

func (m *RepeaterCommand) Execute() error {
	for {
		err := m.command.Execute()
		if err != nil {
			return err
		}
	}
}
