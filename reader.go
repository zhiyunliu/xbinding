package xbinding

//绑定数据的源reader
type Reader interface {
	ReadObject() (any, error)
}

type BytesReader []byte

func (b BytesReader) ReadObject() (any, error) {
	return []byte(b), nil
}

type MapReader map[string][]string

func (b MapReader) ReadObject() (any, error) {
	return map[string][]string(b), nil
}

type SMapReader map[string]string

func (b SMapReader) ReadObject() (any, error) {
	return map[string]string(b), nil
}

type XMapReader map[string]any

func (b XMapReader) ReadObject() (any, error) {
	return map[string]any(b), nil
}

type ReaderWrapper struct {
	Data any
}

func (r *ReaderWrapper) ReadObject() (any, error) {
	return r.Data, nil
}

type MultipartReqestInfo struct {
	Boundary string
	Body     any
}
