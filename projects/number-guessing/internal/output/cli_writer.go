package output

import "fmt"

type CLIWriter struct {
}

func NewCLIWriter() Writer {
	return &CLIWriter{}
}
func (w *CLIWriter) Write(value string) {
	_, err := fmt.Printf("\033[1;32m%s\033[0m\n\n", value)
	if err != nil {
		panic(err)
	}
}
