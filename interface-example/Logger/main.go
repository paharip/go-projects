package main

import (
	"fmt"
	"os"
	"time"
)

type Logger interface {
	Log(messase string)
}

type ConsoleLogger struct{}

func (ConsoleLogger) Log(message string) {
	fmt.Printf("[Console][%s] %s\n", time.Now().Format(time.RFC822), message)
}

type FileLogger struct {
	Filename string
}

func (f FileLogger) Log(message string) {

	file, err := os.OpenFile(f.Filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("[FileLogger] Error opening file:", err)
		return
	}
	defer file.Close()

	timestamp := time.Now().Format(time.RFC822)
	fmt.Fprintf(file, "[File][%s] %s\n", timestamp, message)

}

func Process(logger Logger) {
	logger.Log("Starting process...")
	logger.Log("Process completed!")
}

func main() {
	Process(ConsoleLogger{})
	Process(FileLogger{})
}
