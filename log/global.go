package log

type Logger interface {
	Info(str string)
}

var ecsGlobalLogger Logger = newFallbackLogger()

func Reset() {
	ecsGlobalLogger = newFallbackLogger()
}

func Set(logger Logger) {
	if logger == nil {
		return
	}

	ecsGlobalLogger = logger
}

func Get() Logger {
	return ecsGlobalLogger
}
