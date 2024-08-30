package kernel

type LoggerInterface interface {
	Boot(lf string)
	LogInfo(message string)
	LogError(err error, message string)
}
