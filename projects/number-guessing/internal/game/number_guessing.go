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

	g.writer.Write("Welcome to the Number Guessing Game!")
	g.writer.Write("Let's set up the game parameters.")
	g.writer.Write("Enter the maximum number to predict")
	g.max = g.reader.ReadInt()

	g.writer.Write("Enter the maximum number of attempts allowed")
	g.maxAttempts = g.reader.ReadInt()
	return nil
}

func (g *NumberGuessingGame) Play() error {

	g.writer.Write("I'm thinking of a number between 1 and " + strconv.Itoa(g.max) + ".")

	g.target = rand.Intn(g.max) + 1
	attempts := 0

	g.writer.Write("You have " + strconv.Itoa(g.maxAttempts) + " chances to guess the correct number.")

	now := time.Now()

	start := 1
	finish := g.max

	for {

		if attempts >= g.maxAttempts {
			g.writer.Write("Maximum attempts reached. Game over!")
			g.writer.Write("The number was: " + strconv.Itoa(g.target) + "")
			break
		}

		attempts++
		g.writer.Write("Attempt " + strconv.Itoa(attempts) + " of " + strconv.Itoa(g.maxAttempts) + ".")
		g.writer.Write("Guess a number between " + strconv.Itoa(start) + " and " + strconv.Itoa(finish) + ":")

		inputGuess := g.reader.ReadInt()

		if inputGuess < 1 || inputGuess > g.max {
			g.writer.Write("Invalid guess. Please try again.")
			continue
		}

		if inputGuess < g.target {
			g.writer.Write("Low! Try again.")
			start = inputGuess
		} else if inputGuess > g.target {
			g.writer.Write("High! Try again.")
			finish = inputGuess
		} else {
			g.writer.Write("Congratulations! You've guessed the number in " + strconv.Itoa(attempts) + " attempts.")
			break
		}
	}

	end := time.Now()
	duration := end.Sub(now)
	g.writer.Write("Game duration: " + duration.String())

	return nil
}
