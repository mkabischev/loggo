package loggo

type IProcessor interface {
	Process(entry *Entry)
}
