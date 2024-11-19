// Copyright 2014 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package binding

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"net/url"

	"github.com/zhiyunliu/golibs/bytesconv"
	"github.com/zhiyunliu/golibs/xreflect"
)

func (formBinding) Marshal(v interface{}) ([]byte, error) {
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

func (formPostBinding) Marshal(v interface{}) ([]byte, error) {
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

func (b formMultipartBinding) Marshal(v interface{}) ([]byte, error) {
	mapVal, err := xreflect.AnyToMap(v, xreflect.WithMaxDepth(1))
	if err != nil {
		return nil, err
	}

	byteBuffer := &bytes.Buffer{}
	writer := multipart.NewWriter(byteBuffer)
	for k, v := range mapVal {
		writer.WriteField(k, fmt.Sprint(v))
	}

	writer.Close()
	return byteBuffer.Bytes(), nil
}
