package output

import "fmt"

type CLIWriter struct {
}

func NewCLIWriter() Writer {
	return &CLIWriter{}
}
func (w *CLIWriter) WriteInt(value int) {
	_, err := fmt.Printf("%d\n", value)
	if err != nil {
		panic(err)
	}
}
func (w *CLIWriter) WriteString(value string) {
	_, err := fmt.Printf("%s\n", value)
	if err != nil {
		panic(err)
	}
}
