package binding

import "encoding/xml"

func (xmlBinding) Marshal(v interface{}) ([]byte, error) {
	return xml.Marshal(v)
}
func (xmlBinding) ContentType() string {
	return MIMEXML
}
