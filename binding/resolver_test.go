package binding

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zhiyunliu/xbinding"
)

func Test_getBinding(t *testing.T) {
	testObj := struct {
		XMLName xml.Name `json:"-" yaml:"-" xml:"cxml"`
		A       string   `json:"a" yaml:"a" xml:"a"`
		B       string   `json:"b" yaml:"b" xml:"b"`
	}{A: "A", B: "B"}

	bindingObj := getBinding(xbinding.NewOptions(xbinding.WithContentType("json")))

	val, err := bindingObj.Marshal(testObj)

	assert.Nil(t, err)
	assert.Equal(t, `{"a":"A","b":"B"}`, string(val))
	//----------------------
	bindingObj = getBinding(xbinding.NewOptions(xbinding.WithContentType("text")))

	val, err = bindingObj.Marshal("B")

	assert.Nil(t, err)
	assert.Equal(t, `B`, string(val))

	//----------------------
	bindingObj = getBinding(xbinding.NewOptions(xbinding.WithContentType("form")))

	val, err = bindingObj.Marshal(testObj)

	assert.Nil(t, err)
	assert.Equal(t, `a=A&b=B`, string(val))

	//----------------------
	bindingObj = getBinding(xbinding.NewOptions(xbinding.WithContentType("form-urlencoded")))

	val, err = bindingObj.Marshal(testObj)

	assert.Nil(t, err)
	assert.Equal(t, `a=A&b=B`, string(val))

	//----------------------
	bindingObj = getBinding(xbinding.NewOptions(xbinding.WithContentType("xml")))

	val, err = bindingObj.Marshal(testObj)

	assert.Nil(t, err)
	assert.Equal(t, `<cxml><a>A</a><b>B</b></cxml>`, string(val))

	//----------------------
	bindingObj = getBinding(xbinding.NewOptions(xbinding.WithContentType("yaml")))

	val, err = bindingObj.Marshal(testObj)

	assert.Nil(t, err)
	assert.Equal(t, "a: A\nb: B\n", string(val))
}
