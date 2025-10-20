package main

import (
	"fmt"

	"github.com/ppahari/go-project/factory-functions/domain"
)

func main() {
	cfg := domain.NewConfig("localhost", 8080)
	fmt.Println(cfg.Host, cfg.Port)
}
