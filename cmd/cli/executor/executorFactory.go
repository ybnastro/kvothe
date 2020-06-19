package executor

//KvotheCLIExecutor is the interface that defines the behavior of the command executor in the CLI App
type KvotheCLIExecutor interface {
	RunCommand(args []string) error
}

//GetKvotheCLIExecutor gets the executor based on command name
func GetKvotheCLIExecutor(command string) KvotheCLIExecutor {
	switch command {
	case "ping":
		return new(PingExecutor)
	case "greet":
		return new(GreetExecutor)
	default:
		return nil
	}
}
