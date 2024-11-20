// Copyright 2017 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package binding

import (
	"fmt"
	"net/url"

	"github.com/zhiyunliu/golibs/bytesconv"
	"github.com/zhiyunliu/golibs/xreflect"
)

func (queryBinding) Marshal(v interface{}) ([]byte, error) {
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

func (queryBinding) ContentType() string {
	return MIMEPOSTForm
}
