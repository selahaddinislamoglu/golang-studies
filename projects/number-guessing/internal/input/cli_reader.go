package input

import "fmt"

type CLIReader struct {
}

func NewCLIReader() Reader {
	return &CLIReader{}
}
func (r *CLIReader) ReadInt() int {
	var value int
	_, err := fmt.Scan(&value)
	if err != nil {
		panic(err)
	}
	return value
}
func (r *CLIReader) ReadString() string {
	var value string
	_, err := fmt.Scan(&value)
	if err != nil {
		panic(err)
	}
	return value
}
