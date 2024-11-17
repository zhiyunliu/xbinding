// Copyright 2014 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package binding

import (
	"fmt"
	"mime/multipart"
	"net/url"

	"github.com/zhiyunliu/xbinding"
)

const defaultMemory = 32 << 20

type formBinding struct{}
type formPostBinding struct{}
type formMultipartBinding struct {
	params map[string]string
}

func (formBinding) Name() string {
	return "form"
}

func (formBinding) Bind(reader xbinding.Reader, obj interface{}) error {
	dataObj, err := reader.ReadObject()
	if err != nil {
		return err
	}
	vs, ok := dataObj.(url.Values)
	if !ok {
		return fmt.Errorf("form-urlencoded binding requires url.Values object")
	}

	if err := mapForm(obj, vs); err != nil {
		return err
	}
	return validate(obj)
}

func (formPostBinding) Name() string {
	return "form-urlencoded"
}

func (formPostBinding) Bind(reader xbinding.Reader, obj interface{}) error {
	dataObj, err := reader.ReadObject()
	if err != nil {
		return err
	}
	vs, ok := dataObj.(url.Values)
	if !ok {
		return fmt.Errorf("form-urlencoded binding requires url.Values object")
	}
	if err := mapForm(obj, vs); err != nil {
		return err
	}
	return validate(obj)
}

func (b formMultipartBinding) Name() string {
	return "multipart/form-data"
}

func (b formMultipartBinding) Bind(reader xbinding.Reader, obj interface{}) error {

	dataObj, err := reader.ReadObject()
	if err != nil {
		return err
	}
	reqInfo, ok := dataObj.(*multipartReqInfo)
	if !ok {
		return fmt.Errorf("multipart/form-data binding requires *multipartReqInfo object")

	}

	multiPartReader := multipart.NewReader(reqInfo.Body, reqInfo.Boundary)

	multiForm, err := multiPartReader.ReadForm(defaultMemory)
	if err != nil {
		return err
	}

	if err := mappingByPtr(obj, (*multipartRequest)(multiForm), "form"); err != nil {
		return err
	}

	return validate(obj)
}
