package interfaces

type ILogger interface {
	LogInfo(payload string)
	Close()
}
