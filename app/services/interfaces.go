package services

import "github.com/gorilla/sessions"

type ConfigInterface interface {
	InitConfig()
	Get(key string) string
	GetOrPanic(key string) string
}

type LoggerInterface interface {
	Boot(lf string)
	LogInfo(message string)
	LogError(err error, message string)
}

type StoreInterface sessions.Store

type StorageInterface interface {
	Put(filePath string, content []byte) error
	Get(filePath string) ([]byte, error)
	Delete(filePath string) error
	Exists(filePath string) (bool, error)
}
