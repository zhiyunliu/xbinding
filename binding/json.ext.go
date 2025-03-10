// Copyright 2014 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package binding

import "encoding/json"

func (jsonBinding) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (jsonBinding) ContentType() string {
	return MIMEJSON
}
