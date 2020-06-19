package executor

import (
	"fmt"
)

type PingExecutor struct{}

func (cp *PingExecutor) RunCommand(args []string) error {
	fmt.Println("[Ping Executor] PONG!")
	return nil
}
