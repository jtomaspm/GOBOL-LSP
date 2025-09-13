package application

import (
	"log"
	"os"
)

type Logger struct {
	logger *log.Logger
}

func NewLogger(settings Settings) *Logger {
	filepath := settings.LogPath
	if filepath != "" {
		return &Logger{logger: getLogger(filepath)}
	}
	return &Logger{logger: nil}
}

func (l *Logger) Printf(format string, v ...any) {
	if l.logger != nil {
		l.logger.Printf(format, v...)
	}
}

func (l *Logger) Println(v ...any) {
	if l.logger != nil {
		l.logger.Println(v...)
	}
}

func getLogger(filepath string) *log.Logger {
	f, err := os.OpenFile(filepath, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	return log.New(f, "[GOBOL-LSP]", log.Ldate|log.Ltime|log.Lshortfile)
}
