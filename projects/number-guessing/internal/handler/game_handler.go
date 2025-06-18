package handler

import (
	"github.com/selahaddinislamoglu/golang-studies/projects/number-guessing/internal/game"
	"github.com/selahaddinislamoglu/golang-studies/projects/number-guessing/internal/input"
	"github.com/selahaddinislamoglu/golang-studies/projects/number-guessing/internal/output"
)

type GameHandler struct {
	game   game.Game
	reader input.Reader
	writer output.Writer
}

func NewGameHandler(game game.Game, reader input.Reader, writer output.Writer) Handler {
	return &GameHandler{
		game:   game,
		reader: reader,
		writer: writer,
	}
}
func (h *GameHandler) Run() error {
	if err := h.game.Initialize(); err != nil {
		return err
	}
	for {
		if err := h.game.Play(); err != nil {
			return err
		}
		h.writer.WriteString("Play again? (Y/N)> ")
		input := h.reader.ReadString()
		if input != "Y" && input != "y" {
			h.writer.WriteString("Thanks for playing!")
			break
		}
	}
	return nil
}
