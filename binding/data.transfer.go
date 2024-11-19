package binding

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/zhiyunliu/golibs/xtypes"
)

func transferMapArrayData(dataObj any) (sourceData map[string][]string, err error) {

	switch tmp := dataObj.(type) {
	case url.Values:
		sourceData = tmp
	case map[string][]string:
		sourceData = tmp
	case http.Header:
		sourceData = tmp
	case map[string]string:
		sourceData = map[string][]string{}
		for k, v := range tmp {
			sourceData[k] = []string{v}
		}
	case xtypes.SMap:
		sourceData = map[string][]string{}
		for k, v := range tmp {
			sourceData[k] = []string{v}
		}
	case map[string]any:
		sourceData = map[string][]string{}
		for k, v := range tmp {
			sourceData[k] = []string{fmt.Sprint(v)}
		}
	case xtypes.XMap:
		sourceData = map[string][]string{}
		for k, v := range tmp {
			sourceData[k] = []string{fmt.Sprint(v)}
		}
	default:
		err = fmt.Errorf("binding datatype error[%T]", dataObj)
	}
	return
}

func transferIoReader(dataObj any) (reader io.Reader, err error) {
	switch tmp := dataObj.(type) {
	case io.Reader:
		reader = tmp
	case []byte:
		reader = bytes.NewReader(tmp)
	default:
		err = fmt.Errorf("binding datatype error[%T]", dataObj)
	}
	return
}
