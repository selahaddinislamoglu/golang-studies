package game

type Game interface {
	Initialize() error
	Play() error
}
