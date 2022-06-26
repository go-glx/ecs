package log

import "log"

type fallbackLogger struct{}

func newFallbackLogger() *fallbackLogger {
	return &fallbackLogger{}
}

func (d fallbackLogger) Info(str string) {
	log.Printf("[info] %s\n", str)
}
