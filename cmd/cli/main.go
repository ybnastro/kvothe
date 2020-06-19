package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/SurgicalSteel/kvothe/cmd/cli/executor"
)

func main() {
	registerHandlers()
	arguments := os.Args
	if len(arguments) > 1 {
		command := strings.ToLower(arguments[1])
		err := validateCommand(command)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		params := arguments[2:]

		handler := getHandler(command)
		err = handler.validate(params)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		commandExecutor := executor.GetKvotheCLIExecutor(command)
		if commandExecutor != nil {
			err = commandExecutor.RunCommand(params)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
		} else {
			fmt.Printf("Invalid command, command executor not found for command %s\n", command)
			return
		}
	} else {
		showHelp()
	}

}
