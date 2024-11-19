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

	bindingObj := getBinding(xbinding.NewOptions(xbinding.WithContextType("json")))
	m, ok := bindingObj.(Marshaler)
	assert.True(t, ok)

	val, err := m.Marshal(testObj)

	assert.Nil(t, err)
	assert.Equal(t, `{"a":"A","b":"B"}`, string(val))
	//----------------------
	bindingObj = getBinding(xbinding.NewOptions(xbinding.WithContextType("text")))
	m, ok = bindingObj.(Marshaler)
	assert.True(t, ok)

	val, err = m.Marshal("B")

	assert.Nil(t, err)
	assert.Equal(t, `B`, string(val))

	//----------------------
	bindingObj = getBinding(xbinding.NewOptions(xbinding.WithContextType("form")))
	m, ok = bindingObj.(Marshaler)
	assert.True(t, ok)

	val, err = m.Marshal(testObj)

	assert.Nil(t, err)
	assert.Equal(t, `a=A&b=B`, string(val))

	//----------------------
	bindingObj = getBinding(xbinding.NewOptions(xbinding.WithContextType("form-urlencoded")))
	m, ok = bindingObj.(Marshaler)
	assert.True(t, ok)

	val, err = m.Marshal(testObj)

	assert.Nil(t, err)
	assert.Equal(t, `a=A&b=B`, string(val))

	//----------------------
	bindingObj = getBinding(xbinding.NewOptions(xbinding.WithContextType("xml")))
	m, ok = bindingObj.(Marshaler)
	assert.True(t, ok)

	val, err = m.Marshal(testObj)

	assert.Nil(t, err)
	assert.Equal(t, `<cxml><a>A</a><b>B</b></cxml>`, string(val))

	//----------------------
	bindingObj = getBinding(xbinding.NewOptions(xbinding.WithContextType("yaml")))
	m, ok = bindingObj.(Marshaler)
	assert.True(t, ok)

	val, err = m.Marshal(testObj)

	assert.Nil(t, err)
	assert.Equal(t, "a: A\nb: B\n", string(val))
}
