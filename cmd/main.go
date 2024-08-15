package main

import (
	"log"
	"orangepirepo/internal/service"
)

func main() {
	s := service.NewServer("0.0.0.0:9999")
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
