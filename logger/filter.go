package logger

type Filter interface {
	DoFilter(any) any
	IsFilter(string, any) bool
}
