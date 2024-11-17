package xbinding

type Reader interface {
	ReadObject() (any, error)
}
