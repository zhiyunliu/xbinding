// Copyright 2018 Gin Core Team.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package binding

import "github.com/zhiyunliu/xbinding"

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
