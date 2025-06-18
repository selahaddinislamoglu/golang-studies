package input

type Reader interface {
	ReadInt() int
	ReadString() string
}
