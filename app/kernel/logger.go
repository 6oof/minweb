package kernel

import (
	"log"
	"os"
	"sync"

	"github.com/6oof/minweb/app/configs"
)

type AppLogger struct {
	logger *log.Logger
	mu     sync.Mutex
	once   sync.Once
}

// Boot initializes the logger by setting up the log file.
func (l *AppLogger) Boot() {
	lc := configs.LoggerConfig()
	l.once.Do(func() {
		logFile, err := os.OpenFile(lc.File, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}

		l.logger = log.New(logFile, "", log.Ldate|log.Ltime)
	})
}

// LogInfo writes an info-level log message to the log file.
func (l *AppLogger) LogInfo(message string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.logger.Println("INFO:", message)
}

// LogError writes an error-level log message to the log file.
func (l *AppLogger) LogError(err error, message string) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.logger.Println("ERROR:", message, "-", err)
}
