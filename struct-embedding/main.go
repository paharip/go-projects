package main

import "fmt"

type Logger struct{}

func (l Logger) Log(msg string) {
	fmt.Println("[LOG]:", msg)
}

//composed struct
type Server struct {
	Logger
	Host string
	Port int
}

func (s Server) Start() {
	s.Log(fmt.Sprintf("Starting server on %s:%d", s.Host, s.Port))
}

func main() {
	s := Server{
		Logger: Logger{},
		Host:   "localhost",
		Port:   8080,
	}
	s.Start()
}
