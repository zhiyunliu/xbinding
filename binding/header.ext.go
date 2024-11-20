package binding

import (
	"fmt"
	"net/url"

	"github.com/zhiyunliu/golibs/bytesconv"
	"github.com/zhiyunliu/golibs/xreflect"
)

func (b headerBinding) Marshal(v interface{}) ([]byte, error) {
	mapVal, err := xreflect.AnyToMap(v)
	if err != nil {
		return nil, err
	}
	vals := url.Values{}

	for k, v := range mapVal {
		vals.Set(k, fmt.Sprint(v))
	}
	return bytesconv.StringToBytes(vals.Encode()), nil
}

func (headerBinding) ContentType() string {
	return MIMEPOSTForm
}
