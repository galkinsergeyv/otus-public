package macrocommands

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMove_MoveWhileFuelExists(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	movable := NewMockIMovable(ctrl)
	movable.EXPECT().GetPosition().Return(NewCoords(12, 5))
	movable.EXPECT().GetVelocity().Return(NewCoords(-7, 3))
	movable.EXPECT().SetPosition(gomock.Any())

	fuelBurnable := NewMockIFuelBurnable(ctrl)
	fuelBurnable.EXPECT().GetLevel().Return(1).Times(2)
	fuelBurnable.EXPECT().SetLevel(0)
	fuelBurnable.EXPECT().GetLevel().Return(0)

	commands := []ICommand{
		NewCheckFuelCommand(fuelBurnable),
		NewBurnFuelCommand(fuelBurnable),
		NewMoveCommand(movable),
	}
	macro := NewMacroCommand(commands)

	move := NewRepeaterCommand(macro)

	// act
	err := move.Execute()

	// assert
	assert.Error(t, err)
	assert.Equal(t, NoFuelError, err)
}
