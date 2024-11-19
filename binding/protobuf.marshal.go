// Copyright 2014 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package binding

import (
	"fmt"

	"google.golang.org/protobuf/proto"
)

func (protobufBinding) Marshal(v interface{}) ([]byte, error) {
	msg, ok := v.(proto.Message)
	if !ok {
		return nil, fmt.Errorf("obj is not ProtoMessage,%T", v)
	}
	return proto.Marshal(msg)
}
