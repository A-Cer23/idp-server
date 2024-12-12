package main

import (
	"fmt"
	"log"
)

type Logger struct {
}

var logger = Logger{}

func (l *Logger) Info(message string) {
	l.log("INFO", message)
}

func (l *Logger) Error(message string, err error) {
	l.log("ERROR", fmt.Sprintf("%v: %v", message, err))
}

func (l *Logger) log(level string, message string) {
	log.Printf("[%s] %v", level, message)
}
