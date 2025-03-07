// Copyright 2018 Gin Core Team.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package binding

import (
	"fmt"
	"net/url"

	"github.com/zhiyunliu/golibs/bytesconv"
	"github.com/zhiyunliu/golibs/xreflect"
	"github.com/zhiyunliu/xbinding"
)

type uriBinding struct{}

func (uriBinding) Name() string {
	return "uri"
}

func (uriBinding) Bind(reader xbinding.Reader, obj interface{}) error {

	dataObj, err := reader.ReadObject()
	if err != nil {
		return err
	}
	realData, err := transferMapArrayData(dataObj)
	if err != nil {
		return err
	}

	if err := mapURI(obj, realData); err != nil {
		return err
	}
	return validate(obj)
}

func (uriBinding) ContentType() string {
	return MIMEPOSTForm
}

func (uriBinding) Marshal(v interface{}) ([]byte, error) {
	mapVal, err := xreflect.AnyToMap(v, xreflect.WithMaxDepth(1))
	if err != nil {
		return nil, err
	}
	vals := url.Values{}

	for k, v := range mapVal {
		vals.Add(k, fmt.Sprint(v))
	}
	return bytesconv.StringToBytes(vals.Encode()), nil
}
