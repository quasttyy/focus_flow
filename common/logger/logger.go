package logger

import (
	"log"
	"os"
)

type Logger struct {
	info  *log.Logger
	warn  *log.Logger
	err   *log.Logger
	debug *log.Logger
}

var Log *Logger // глобальная ссылка (по аналогии с zap.L())

// Init — инициализация глобального логгера
func Init() {
	Log = &Logger{
		info:  log.New(os.Stdout, "INFO:  ", log.Ldate|log.Ltime|log.Lshortfile),
		warn:  log.New(os.Stdout, "WARN:  ", log.Ldate|log.Ltime|log.Lshortfile),
		err:   log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
		debug: log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile),
	}

	Log.Info("logger initialized")
}

// Info — обычные инфо-сообщения
func (l *Logger) Info(v ...any) {
	l.info.Println(v...)
}

// Warn — предупреждения
func (l *Logger) Warn(v ...any) {
	l.warn.Println(v...)
}

// Error — ошибки
func (l *Logger) Error(v ...any) {
	l.err.Println(v...)
}

// Debug — отладочная информация
func (l *Logger) Debug(v ...any) {
	// можно выключать в проде через ENV
	if os.Getenv("APP_ENV") == "prod" {
		return
	}
	l.debug.Println(v...)
}

// Fatal — критические ошибки (завершает выполнение)
func (l *Logger) Fatal(v ...any) {
	l.err.Println(v...)
	os.Exit(1)
}
