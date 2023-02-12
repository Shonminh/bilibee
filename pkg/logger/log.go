package logger


var logger Logger

// LogDebug logs a message at level Debug on the standard logger.Log
func LogDebug(args ...interface{}) {
	logger.LogDebug(args...)
}

// LogPrint logs a message at level Info on the standard logger.Log
func LogPrint(args ...interface{}) {
	logger.LogPrint(args...)
}

// LogInfo logs a message at level Info on the standard logger.Log
func LogInfo(args ...interface{}) {
	logger.LogInfo(args...)
}

// LogWarn logs a message at level Warn on the standard logger.Log
func LogWarn(args ...interface{}) {
	logger.LogWarn(args...)
}

// LogWarning logs a message at level Warn on the standard logger.Log
func LogWarning(args ...interface{}) {
	logger.LogWarning(args...)
}

// LogError logs a message at level Error on the standard logger.Log
func LogError(args ...interface{}) {
	logger.LogError(args...)
}

// LogPanic logs a message at level Panic on the standard logger.Log
func LogPanic(args ...interface{}) {
	logger.LogPanic(args...)
}

// LogFatal logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func LogFatal(args ...interface{}) {
	logger.LogFatal(args...)
}

// LogTracef logs a message at level Trace on the standard logger.Log
func LogTracef(format string, args ...interface{}) {
	logger.LogTracef(format, args...)
}

// LogDebugf logs a message at level Debug on the standard logger.Log
func LogDebugf(format string, args ...interface{}) {
	logger.LogDebugf(format, args...)
}

// LogPrintf logs a message at level Info on the standard logger.Log
func LogPrintf(format string, args ...interface{}) {
	logger.LogPrintf(format, args...)
}

// Printf logs a message at level Info on the standard logger.Log
func Printf(format string, args ...interface{}) {
	logger.LogPrintf(format, args...)
}

// LogInfof logs a message at level Info on the standard logger.Log
func LogInfof(format string, args ...interface{}) {
	logger.LogInfof(format, args...)
}

// LogWarnf logs a message at level Warn on the standard logger.Log
func LogWarnf(format string, args ...interface{}) {
	logger.LogWarnf(format, args...)
}

// LogWarningf logs a message at level Warn on the standard logger.Log
func LogWarningf(format string, args ...interface{}) {
	logger.LogWarningf(format, args...)
}

// LogErrorf logs a message at level Error on the standard logger.Log
func LogErrorf(format string, args ...interface{}) {
	logger.LogErrorf(format, args...)
}

// LogPanicf logs a message at level Panic on the standard logger.Log
func LogPanicf(format string, args ...interface{}) {
	logger.LogPanicf(format, args...)
}

// LogFatalf logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func LogFatalf(format string, args ...interface{}) {
	logger.LogFatalf(format, args...)
}

// LogTraceln logs a message at level Trace on the standard logger.Log
func LogTraceln(args ...interface{}) {
	logger.LogTraceln(args...)
}

// LogDebugln logs a message at level Debug on the standard logger.Log
func LogDebugln(args ...interface{}) {
	logger.LogDebugln(args...)
}

// LogPrintln logs a message at level Info on the standard logger.Log
func LogPrintln(args ...interface{}) {
	logger.LogPrintln(args...)
}

// LogInfoln logs a message at level Info on the standard logger.Log
func LogInfoln(args ...interface{}) {
	logger.LogInfoln(args...)
}

// LogWarnln logs a message at level Warn on the standard logger.Log
func LogWarnln(args ...interface{}) {
	logger.LogWarnln(args...)
}

// LogWarningln logs a message at level Warn on the standard logger.Log
func LogWarningln(args ...interface{}) {
	logger.LogWarningln(args...)
}

// LogErrorln logs a message at level Error on the standard logger.Log
func LogErrorln(args ...interface{}) {
	logger.LogErrorln(args...)
}

// LogPanicln logs a message at level Panic on the standard logger.Log
func LogPanicln(args ...interface{}) {
	logger.LogPanicln(args...)
}

// LogFatalln logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func LogFatalln(args ...interface{}) {
	logger.LogFatalln(args...)
}

