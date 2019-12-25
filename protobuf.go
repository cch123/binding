// Copyright 2014 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package binding

import (
	"github.com/cch123/binding/internal/proto"

	"io/ioutil"
	"net/http"
)

type protobufBinding struct{}

func (protobufBinding) Name() string {
	return "protobuf"
}

func (protobufBinding) Bind(req *http.Request, obj interface{}) error {

	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}

	if err = proto.Unmarshal(buf, obj.(proto.Message)); err != nil {
		return err
	}

	return nil
}
