package handle_exceptions

import (
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

var SomeError = errors.New("some error")

func TestHandleExceptions_RepeatThenLog_CmdAndLoggerCalledOnce(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := NewMockILogger(ctrl)
	logger.EXPECT().Log(gomock.Any()).Times(1)
	cmd := NewMockICommand(ctrl)
	cmd.EXPECT().GetType().Return("RepeatThenLogCommand").AnyTimes()
	cmd.EXPECT().Execute().Return(SomeError).Times(2)

	var exceptionHandlers = map[string]map[error]IExceptionHandler{
		"RepeatThenLogCommand": {
			SomeError: NewRepeatThenLogExceptionHandler(1, logger),
		},
	}
	player := NewPlayer(exceptionHandlers)

	// act
	err := player.Play(cmd)

	// assert
	assert.NoError(t, err)
}

func TestHandleExceptions_RepeatTwiceThenLog_CmdAndLoggerCalledOnce(t *testing.T) {
	// arrange
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	logger := NewMockILogger(ctrl)
	logger.EXPECT().Log(gomock.Any()).Times(1)
	cmd := NewMockICommand(ctrl)
	cmd.EXPECT().GetType().Return("RepeatThenLogCommand").AnyTimes()
	cmd.EXPECT().Execute().Return(SomeError).Times(3)

	var exceptionHandlers = map[string]map[error]IExceptionHandler{
		"RepeatThenLogCommand": {
			SomeError: NewRepeatThenLogExceptionHandler(2, logger),
		},
	}
	player := NewPlayer(exceptionHandlers)

	// act
	err := player.Play(cmd)

	// assert
	assert.NoError(t, err)
}
