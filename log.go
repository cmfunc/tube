package tube

var logger Logger

type Logger interface {
	Info(args ...interface{})
	Infof(format string, args ...interface{})
}

func SetLogger(l Logger)  {logger = l}
