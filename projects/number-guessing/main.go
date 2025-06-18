package main

import (
	"github.com/selahaddinislamoglu/golang-studies/projects/number-guessing/cmd"
	"github.com/selahaddinislamoglu/golang-studies/projects/number-guessing/internal/game"
	"github.com/selahaddinislamoglu/golang-studies/projects/number-guessing/internal/handler"
	"github.com/selahaddinislamoglu/golang-studies/projects/number-guessing/internal/input"
	"github.com/selahaddinislamoglu/golang-studies/projects/number-guessing/internal/output"
)

func main() {

	reader := input.NewCLIReader()
	writer := output.NewCLIWriter()
	game := game.NewNumberGuessingGame(reader, writer)
	handler := handler.NewGameHandler(game, reader, writer)
	command := cmd.NewCommand(handler)
	if err := command.Execute(); err != nil {
		panic(err)
	}
}
