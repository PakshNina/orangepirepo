package main

import (
	"log"
	_ "net/http/pprof"
	"orangepirepo/internal/board"
	"orangepirepo/internal/service"
)

func main() {
	b := board.NewBoard()
	s := service.NewServer("0.0.0.0:9999", b)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
