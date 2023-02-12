package logger

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetFormatter(&LogFormatter{})
	logrus.SetLevel(logrus.DebugLevel)
	file, err := GetOutput()
	if err == nil {
		logrus.SetOutput(file)
	} else {
		logrus.Info("Failed to log to file, using default stderr")
	}
}

type Logger struct {
}

func GetOutput() (io.Writer, error) {
	file, err := os.OpenFile("log/data.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	return file, err
}

// LogDebug logs a message at level Debug on the standard logger.
func (l Logger) LogDebug(args ...interface{}) {
	logrus.Debug(args...)
}

// LogPrint logs a message at level Info on the standard logger.
func (l Logger) LogPrint(args ...interface{}) {
	logrus.Print(args...)
}

// LogInfo logs a message at level Info on the standard logger.
func (l Logger) LogInfo(args ...interface{}) {
	logrus.Info(args...)
}

// LogWarn logs a message at level Warn on the standard logger.
func (l Logger) LogWarn(args ...interface{}) {
	logrus.Warn(args...)
}

// LogWarning logs a message at level Warn on the standard logger.
func (l Logger) LogWarning(args ...interface{}) {
	logrus.Warning(args...)
}

// LogError logs a message at level Error on the standard logger.
func (l Logger) LogError(args ...interface{}) {
	logrus.Error(args...)
}

// LogPanic logs a message at level Panic on the standard logger.
func (l Logger) LogPanic(args ...interface{}) {
	logrus.Panic(args...)
}

// LogFatal logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func (l Logger) LogFatal(args ...interface{}) {
	logrus.Fatal(args...)
}

// LogTracef logs a message at level Trace on the standard logger.
func (l Logger) LogTracef(format string, args ...interface{}) {
	logrus.Tracef(format, args...)
}

// LogDebugf logs a message at level Debug on the standard logger.
func (l Logger) LogDebugf(format string, args ...interface{}) {
	logrus.Debugf(format, args...)
}

// LogPrintf logs a message at level Info on the standard logger.
func (l Logger) LogPrintf(format string, args ...interface{}) {
	logrus.Printf(format, args...)
}

// Printf logs a message at level Info on the standard logger.
func (l Logger) Printf(format string, args ...interface{}) {
	logrus.Printf(format, args...)
}

// LogInfof logs a message at level Info on the standard logger.
func (l Logger) LogInfof(format string, args ...interface{}) {
	logrus.Infof(format, args...)
}

// LogWarnf logs a message at level Warn on the standard logger.
func (l Logger) LogWarnf(format string, args ...interface{}) {
	logrus.Warnf(format, args...)
}

// LogWarningf logs a message at level Warn on the standard logger.
func (l Logger) LogWarningf(format string, args ...interface{}) {
	logrus.Warningf(format, args...)
}

// LogErrorf logs a message at level Error on the standard logger.
func (l Logger) LogErrorf(format string, args ...interface{}) {
	logrus.Errorf(format, args...)
}

// LogPanicf logs a message at level Panic on the standard logger.
func (l Logger) LogPanicf(format string, args ...interface{}) {
	logrus.Panicf(format, args...)
}

// LogFatalf logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func (l Logger) LogFatalf(format string, args ...interface{}) {
	logrus.Fatalf(format, args...)
}

// LogTraceln logs a message at level Trace on the standard logger.
func (l Logger) LogTraceln(args ...interface{}) {
	logrus.Traceln(args...)
}

// LogDebugln logs a message at level Debug on the standard logger.
func (l Logger) LogDebugln(args ...interface{}) {
	logrus.Debugln(args...)
}

// LogPrintln logs a message at level Info on the standard logger.
func (l Logger) LogPrintln(args ...interface{}) {
	logrus.Println(args...)
}

// LogInfoln logs a message at level Info on the standard logger.
func (l Logger) LogInfoln(args ...interface{}) {
	logrus.Infoln(args...)
}

// LogWarnln logs a message at level Warn on the standard logger.
func (l Logger) LogWarnln(args ...interface{}) {
	logrus.Warnln(args...)
}

// LogWarningln logs a message at level Warn on the standard logger.
func (l Logger) LogWarningln(args ...interface{}) {
	logrus.Warningln(args...)
}

// LogErrorln logs a message at level Error on the standard logger.
func (l Logger) LogErrorln(args ...interface{}) {
	logrus.Errorln(args...)
}

// LogPanicln logs a message at level Panic on the standard logger.
func (l Logger) LogPanicln(args ...interface{}) {
	logrus.Panicln(args...)
}

// LogFatalln logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func (l Logger) LogFatalln(args ...interface{}) {
	logrus.Fatalln(args...)
}
