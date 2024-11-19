// Copyright 2022 Gin Core Team. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package binding

import (
	"io"

	"github.com/pelletier/go-toml/v2"
	"github.com/zhiyunliu/xbinding"
)

type tomlBinding struct{}

func (tomlBinding) Name() string {
	return "toml"
}

func decodeToml(r io.Reader, obj any) error {
	decoder := toml.NewDecoder(r)
	if err := decoder.Decode(obj); err != nil {
		return err
	}
	return decoder.Decode(obj)
}

func (tomlBinding) Marshal(v interface{}) ([]byte, error) {
	return toml.Marshal(v)
}

func (tomlBinding) Bind(reader xbinding.Reader, obj interface{}) error {
	dataObj, err := reader.ReadObject()
	if err != nil {
		return err
	}
	readData, err := transferIoReader(dataObj)
	if err != nil {
		return err
	}

	return decodeToml(readData, obj)
}
