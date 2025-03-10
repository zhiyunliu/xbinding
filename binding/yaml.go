// Copyright 2018 Gin Core Team.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package binding

import (
	"bytes"
	"io"

	"github.com/zhiyunliu/xbinding"
	"gopkg.in/yaml.v3"
)

type yamlBinding struct{}

func (yamlBinding) Name() string {
	return "yaml"
}

func (yamlBinding) Bind(reader xbinding.Reader, obj interface{}) error {

	dataObj, err := reader.ReadObject()
	if err != nil {
		return err
	}
	readData, err := transferIoReader(dataObj)
	if err != nil {
		return err
	}

	return decodeYAML(readData, obj)
}

func (yamlBinding) BindBody(body []byte, obj interface{}) error {
	return decodeYAML(bytes.NewReader(body), obj)
}

func decodeYAML(r io.Reader, obj interface{}) error {
	decoder := yaml.NewDecoder(r)
	if err := decoder.Decode(obj); err != nil {
		return err
	}
	return validate(obj)
}
