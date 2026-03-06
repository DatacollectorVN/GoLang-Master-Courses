package main

import "fmt"

type LogLevel int

const (
	LevelTrace LogLevel = iota
	LevelDebug
	LevelInfo
	LevelWarning
	LevelError
)

var levelNames = []string{"TRACE", "DEBUG", "INFO", "WARNING", "ERROR", "FATAL"}

func (l LogLevel) String() string {
	if l < LevelTrace || l > LevelError {
		return "UNKNOWN"
	}
	return levelNames[l]
}

func printLog(level LogLevel, message string) {
	fmt.Printf("[%s] %s\n", level, message)
}

func main() {
	printLog(LevelTrace, "This is a trace message")
	printLog(LevelDebug, "This is a debug message")
	printLog(LevelInfo, "This is a info message")
	printLog(LevelWarning, "This is a warning message")
	printLog(LevelError, "This is a error message")
}
