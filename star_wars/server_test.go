package star_wars

import (
	"github.com/golang/mock/gomock"
	"testing"
)

func TestMove_DirectMove_PositionChanged(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	movable := NewMockIMovable(ctrl)
	movable.EXPECT().GetPosition().Return(NewCoords(12, 5))
	movable.EXPECT().GetVelocity().Return(NewCoords(-7, 3))
	expectedCoords := NewCoords(5, 8)
	movable.EXPECT().SetPosition(expectedCoords)

	move := NewMove(movable)

	// act
	move.Execute()
}

func TestRotate_ExecuteRotate_VelocityChanged(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	rotatable := NewMockIRotatable(ctrl)
	rotatable.EXPECT().GetPosition().Return(NewCoords(12, 5))
	rotatable.EXPECT().GetAngle().Return(90)
	expectedCoords := NewCoords(-4, -12)
	rotatable.EXPECT().SetVelocity(expectedCoords)

	move := NewRotate(rotatable)

	// act
	move.Execute()
}
