// Copyright 2017 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package binding

import (
	"fmt"
	"net/url"

	"github.com/zhiyunliu/xbinding"
)

type queryBinding struct{}

func (queryBinding) Name() string {
	return "query"
}

func (queryBinding) Bind(reader xbinding.Reader, obj interface{}) error {
	dataObj, err := reader.ReadObject()
	if err != nil {
		return err
	}
	realData, ok := dataObj.(url.Values)
	if !ok {
		return fmt.Errorf("query binding requires an url.Values object")
	}

	if err := mapForm(obj, realData); err != nil {
		return err
	}
	return validate(obj)
}
