package executor

import (
	"errors"
	"fmt"
)

type GreetExecutor struct{}

func (cg *GreetExecutor) RunCommand(args []string) error {
	switch len(args) {
	case 0:
		return errors.New("[Greet Executor] Name undefined")
	case 1:
		fmt.Printf("[Greet Executor] Hello %s, it's nice to see you!\n", args[0])
		return nil
	default:
		names := args[0]
		for i := 1; i < len(args); i++ {
			names = names + "," + args[i]
		}
		fmt.Printf("[Greet Executor] Hello %s, it's nice to see all of you!\n", names)
		return nil
	}
}
