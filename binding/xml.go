// Copyright 2014 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package binding

import (
	"bytes"
	"encoding/xml"
	"io"

	"github.com/zhiyunliu/xbinding"
)

type xmlBinding struct{}

func (xmlBinding) Name() string {
	return "xml"
}

func (xmlBinding) Bind(reader xbinding.Reader, obj interface{}) error {
	dataObj, err := reader.ReadObject()
	if err != nil {
		return err
	}
	readData, err := transferIoReader(dataObj)
	if err != nil {
		return err
	}

	return decodeXML(readData, obj)
}

func (xmlBinding) BindBody(body []byte, obj interface{}) error {
	return decodeXML(bytes.NewReader(body), obj)
}
func decodeXML(r io.Reader, obj interface{}) error {
	decoder := xml.NewDecoder(r)
	if err := decoder.Decode(obj); err != nil {
		return err
	}
	return validate(obj)
}
