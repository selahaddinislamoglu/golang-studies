package game

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/selahaddinislamoglu/golang-studies/projects/number-guessing/internal/input"
	"github.com/selahaddinislamoglu/golang-studies/projects/number-guessing/internal/output"
)

type NumberGuessingGame struct {
	reader      input.Reader
	writer      output.Writer
	min         int
	max         int
	maxAttempts int
	attempts    int
	target      int
}

func NewNumberGuessingGame(reader input.Reader, writer output.Writer) Game {
	return &NumberGuessingGame{
		reader:      reader,
		writer:      writer,
		min:         1,
		max:         100,
		maxAttempts: 10,
		attempts:    0,
		target:      50,
	}
}

func (g *NumberGuessingGame) Initialize() error {

	g.writer.WriteString("Welcome to the Number Guessing Game!")
	g.writer.WriteString("Let's set up the game parameters.")
	g.writer.WriteString("Enter the maximum number to predict")
	g.max = g.reader.ReadInt()

	g.writer.WriteString("Enter the maximum number of attempts allowed")
	g.maxAttempts = g.reader.ReadInt()
	return nil
}

func (g *NumberGuessingGame) Play() error {

	g.writer.WriteString("I'm thinking of a number between 1 and " + strconv.Itoa(g.max) + ".")

	g.target = rand.Intn(g.max) + 1
	attempts := 0

	g.writer.WriteString("You have " + strconv.Itoa(g.maxAttempts) + " chances to guess the correct number.")

	now := time.Now()

	for {

		if attempts >= g.maxAttempts {
			g.writer.WriteString("Maximum attempts reached. Game over!")
			g.writer.WriteString("The number was: " + strconv.Itoa(g.target) + "")
			break
		}

		g.writer.WriteString("Guess a number between 1 and " + strconv.Itoa(g.max) + ":")

		inputGuess := g.reader.ReadInt()

		attempts++

		if inputGuess < 1 || inputGuess > g.max {
			g.writer.WriteString("Invalid guess. Please try again.")
			continue
		}

		if inputGuess < g.target {
			g.writer.WriteString("Low! Try again.")
		} else if inputGuess > g.target {
			g.writer.WriteString("High! Try again.")
		} else {
			g.writer.WriteString("Congratulations! You've guessed the number!")
			break
		}
	}

	end := time.Now()
	duration := end.Sub(now)
	g.writer.WriteString("Game duration: " + duration.String())

	return nil
}
