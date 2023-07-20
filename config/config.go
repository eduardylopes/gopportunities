package config

var (
	l *Logger
)

func Init() error {
	return nil
}

func GetLogger(p string) *Logger {
	l = NewLogger(p)
	return l
}
