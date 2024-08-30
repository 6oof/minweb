package kernel

type LoggerInterface interface {
	Boot()
	LogInfo(message string)
	LogError(err error, message string)
}
