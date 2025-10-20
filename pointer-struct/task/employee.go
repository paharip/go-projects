package task

type Employee struct {
	Name  string
	Level int
}

func Promote(e *Employee) {
	e.Level++
}
