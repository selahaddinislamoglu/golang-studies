package output

type Writer interface {
	WriteInt(value int)
	WriteString(value string)
}
