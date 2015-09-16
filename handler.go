package loggo

// IHandler interface
type IHandler interface {
	Handle(entry *Entry)
}
