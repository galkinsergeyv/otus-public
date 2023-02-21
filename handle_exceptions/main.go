//go:generate mockgen -package=$GOPACKAGE -source=$GOFILE -destination=main_mock.go
package handle_exceptions

import (
	"errors"
	"fmt"
)

var CommandWithoutErrorHandlers = errors.New("command without error handlers")
var UnhandledError = errors.New("unhandled error")

type ILogger interface {
	Log(msg string)
}

type ICommand interface {
	Execute() error
	GetType() string
}

type LogErrorCommand struct {
	cmd    ICommand
	err    error
	logger ILogger
}

func NewLogErrorCommand(cmd ICommand, err error, logger ILogger) *LogErrorCommand {
	return &LogErrorCommand{
		cmd:    cmd,
		err:    err,
		logger: logger,
	}
}

func (c *LogErrorCommand) Execute() error {
	msg := fmt.Errorf("%s throw error: %s", c.cmd.GetType(), c.err.Error())
	c.logger.Log(msg.Error())
	return nil
}

func (c *LogErrorCommand) GetType() string {
	return "LogErrorCommand"
}

type CloseQueueCommand struct {
	queue chan ICommand
}

func NewCloseQueueCommand(queue chan ICommand) *CloseQueueCommand {
	return &CloseQueueCommand{
		queue: queue,
	}
}

func (c *CloseQueueCommand) Execute() error {
	close(c.queue)
	return nil
}

func (c *CloseQueueCommand) GetType() string {
	return "CloseQueueCommand"
}

type LogErrorExceptionHandler struct {
	logger ILogger
}

func NewLogErrorExceptionHandler(logger ILogger) *LogErrorExceptionHandler {
	return &LogErrorExceptionHandler{
		logger: logger,
	}
}

func (h *LogErrorExceptionHandler) Handle(cmd ICommand, err error, queue chan ICommand) error {
	newCmd := NewLogErrorCommand(cmd, err, h.logger)
	queue <- newCmd

	return nil
}

type RequeueCommand struct {
	cmd   ICommand
	queue chan ICommand
}

func NewRequeueCommand(cmd ICommand, queue chan ICommand) *RequeueCommand {
	return &RequeueCommand{
		cmd:   cmd,
		queue: queue,
	}
}

func (c *RequeueCommand) Execute() error {
	c.queue <- c.cmd
	return nil
}

func (c *RequeueCommand) GetType() string {
	return "RequeueCommand"
}

type RequeueExceptionHandler struct{}

func NewRequeueExceptionHandler() *RequeueExceptionHandler {
	return &RequeueExceptionHandler{}
}

func (h *RequeueExceptionHandler) Handle(cmd ICommand, _ error, queue chan ICommand) error {
	newCmd := NewRequeueCommand(cmd, queue)
	queue <- newCmd

	return nil
}

type RepeatThenLogExceptionHandler struct {
	maxRetries int
	counter    map[string]int
	logger     ILogger
}

func NewRepeatThenLogExceptionHandler(maxRetries int, logger ILogger) *RepeatThenLogExceptionHandler {
	counter := map[string]int{}
	return &RepeatThenLogExceptionHandler{
		maxRetries: maxRetries,
		counter:    counter,
		logger:     logger,
	}
}

func (h *RepeatThenLogExceptionHandler) Handle(cmd ICommand, err error, queue chan ICommand) error {
	cmdType := cmd.GetType()
	cmdCount, ok := h.counter[cmdType]
	if !ok || cmdCount < h.maxRetries {
		newCmd := NewRequeueCommand(cmd, queue)
		queue <- newCmd
		if ok {
			h.counter[cmdType] += 1
		} else {
			h.counter[cmdType] = 1
		}

		return nil
	}

	newCmd := NewLogErrorCommand(cmd, err, h.logger)
	queue <- newCmd
	finalCmd := NewCloseQueueCommand(queue)
	queue <- finalCmd
	h.counter[cmdType] += 1

	return nil
}

type IExceptionHandler interface {
	Handle(cmd ICommand, err error, queue chan ICommand) error
}

type ExceptionHandler struct {
	exceptionHandlers map[string]map[error]IExceptionHandler
}

func NewExceptionHandler(exceptionHandlers map[string]map[error]IExceptionHandler) *ExceptionHandler {
	return &ExceptionHandler{
		exceptionHandlers: exceptionHandlers,
	}
}

func (h *ExceptionHandler) Handle(cmd ICommand, err error, queue chan ICommand) error {
	cmdType := cmd.GetType()
	cmdExceptions, ok := h.exceptionHandlers[cmdType]
	if !ok {
		return CommandWithoutErrorHandlers
	}

	handler, ok := cmdExceptions[err]
	if !ok {
		return UnhandledError
	}

	return handler.Handle(cmd, err, queue)
}

type Player struct {
	exceptionHandler IExceptionHandler
}

func NewPlayer(exceptionHandlers map[string]map[error]IExceptionHandler) *Player {
	exceptionHandler := NewExceptionHandler(exceptionHandlers)

	return &Player{
		exceptionHandler: exceptionHandler,
	}
}

func (p *Player) Play(cmd ICommand) error {
	queue := make(chan ICommand, 2)

	queue <- cmd

	for {
		select {
		case nextCmd, ok := <-queue:
			if !ok {
				fmt.Println("Unexpected error with chan")
				return nil
			}

			err := nextCmd.Execute()
			if err != nil {
				handlerErr := p.exceptionHandler.Handle(nextCmd, err, queue)
				if handlerErr != nil {
					return handlerErr
				}
			}
		}
	}
}
