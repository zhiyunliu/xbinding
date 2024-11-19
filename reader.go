package xbinding

//绑定数据的源reader
type Reader interface {
	ReadObject() (any, error)
}
