package main

import (
	"fmt"
)

type Notifier interface {
	Notify(msg string)
}

type BaseNotifier struct{}

func (b BaseNotifier) Notify(msg string) { fmt.Println(msg) }

type LogDecorator struct {
	Notifier
}

func (l LogDecorator) Notify(msg string) {
	fmt.Println("Logging:", msg)
	l.Notifier.Notify(msg)
}

func main() {
	base := BaseNotifier{}
	decorated := LogDecorator{base}
	decorated.Notify("Hello!")

}
