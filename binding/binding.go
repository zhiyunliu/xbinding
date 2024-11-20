// Copyright 2014 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package binding

import (
	"net/http"

	"github.com/zhiyunliu/xbinding"
)

// Content-Type MIME of the most common data formats.
const (
	MIMEJSON              = xbinding.MIMEJSON              // "application/json"
	MIMEHTML              = xbinding.MIMEHTML              // "text/html"
	MIMEXML               = xbinding.MIMEXML               // "application/xml"
	MIMEXML2              = xbinding.MIMEXML2              // "text/xml"
	MIMEPlain             = xbinding.MIMEPlain             // "text/plain"
	MIMEPOSTForm          = xbinding.MIMEPOSTForm          // "application/x-www-form-urlencoded"
	MIMEMultipartPOSTForm = xbinding.MIMEMultipartPOSTForm // "multipart/form-data"
	MIMEPROTOBUF          = xbinding.MIMEPROTOBUF          // "application/x-protobuf"
	MIMEMSGPACK           = xbinding.MIMEMSGPACK           // "application/x-msgpack"
	MIMEMSGPACK2          = xbinding.MIMEMSGPACK2          // "application/msgpack"
	MIMEYAML              = xbinding.MIMEYAML              // "application/x-yaml"
)

// StructValidator is the minimal interface which needs to be implemented in
// order for it to be used as the validator engine for ensuring the correctness
// of the request. Gin provides a default implementation for this using
// https://github.com/go-playground/validator/tree/v10.6.1.
type StructValidator interface {
	// ValidateStruct can receive any kind of type and it should never panic, even if the configuration is not right.
	// If the received type is a slice|array, the validation should be performed travel on every element.
	// If the received type is not a struct or slice|array, any validation should be skipped and nil must be returned.
	// If the received type is a struct or pointer to a struct, the validation should be performed.
	// If the struct is not valid or the validation itself fails, a descriptive error should be returned.
	// Otherwise nil must be returned.
	ValidateStruct(interface{}) error

	// Engine returns the underlying validator engine which powers the
	// StructValidator implementation.
	Engine() interface{}
}

// Validator is the default validator which implements the StructValidator
// interface. It uses https://github.com/go-playground/validator/tree/v10.6.1
// under the hood.
var Validator StructValidator = &defaultValidator{}

// These implement the Binding interface and can be used to bind the data
// present in the request to struct instances.
var (
	JSON          = jsonBinding{}
	XML           = xmlBinding{}
	Form          = formBinding{}
	Query         = queryBinding{}
	FormPost      = formPostBinding{}
	FormMultipart = formMultipartBinding{}
	ProtoBuf      = protobufBinding{}
	YAML          = yamlBinding{}
	Uri           = uriBinding{}
	Header        = headerBinding{}
)

// Default returns the appropriate Binding instance based on the HTTP method
// and the content type.
func Default(method, contentType string) Binding {
	if method == http.MethodGet {
		return Form
	}

	switch contentType {
	case MIMEPlain:
		return Plain
	case MIMEJSON:
		return JSON
	case MIMEXML, MIMEXML2:
		return XML
	case MIMEPROTOBUF:
		return ProtoBuf
	case MIMEYAML:
		return YAML
	case MIMEMultipartPOSTForm:
		return FormMultipart
	default: // case MIMEPOSTForm:
		return nil
	}
}

func validate(obj interface{}) error {
	if Validator == nil {
		return nil
	}
	return Validator.ValidateStruct(obj)
}
