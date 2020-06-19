package main

import (
	"errors"
	"fmt"
)

type handler struct {
	command              string
	mandatoryParamLength int
}

func (h handler) validate(params []string) error {
	if len(params) < h.mandatoryParamLength {
		message := fmt.Sprintf(
			"Invalid parameter length for command %s.\nIt needs %d parameters, but you provided %d parameters",
			h.command,
			h.mandatoryParamLength,
			len(params),
		)
		return errors.New(message)
	}
	return nil
}

var cliHandlers map[string]handler

func registerHandlers() {
	cliHandlers = make(map[string]handler)

	//handle PING
	cliHandlers["ping"] = handler{
		command:              "ping",
		mandatoryParamLength: 0,
	}

	//handle PING
	cliHandlers["greet"] = handler{
		command:              "greet",
		mandatoryParamLength: 1,
	}

}

func validateCommand(command string) error {
	if _, ok := cliHandlers[command]; !ok {
		return errors.New(
			fmt.Sprintf("Command %s is invalid!", command),
		)
	}
	return nil
}

func getHandler(command string) handler {
	return cliHandlers[command]
}
