package binding

import (
	"fmt"

	"github.com/zhiyunliu/golibs/bytesconv"
	"github.com/zhiyunliu/xbinding"
)

type textCodec struct {
}

func (textCodec) Name() string {
	return "text"
}

func (textCodec) Marshal(v interface{}) ([]byte, error) {
	str, _ := v.(string)
	return bytesconv.StringToBytes(str), nil
}

func (textCodec) Bind(reader xbinding.Reader, v interface{}) error {
	dataObj, err := reader.ReadObject()
	if err != nil {
		return err
	}
	str, ok := v.(*string)
	if !ok {
		return fmt.Errorf("text type error,%T", v)
	}

	switch tmp := dataObj.(type) {
	case []byte:
		*str = bytesconv.BytesToString(tmp)
	case *[]byte:
		*str = bytesconv.BytesToString(*tmp)
	case string:
		*str = tmp
	case *string:
		*str = *tmp
	}
	return nil
}
