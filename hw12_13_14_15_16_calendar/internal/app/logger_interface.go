package app

type Logger interface {
	Info(msg string)
	Error(msg string)
	Debug(msg string)
	Print(msg string)
	Warn(msg string)
	Fatal(msg string)
	Panic(msg string)
}
