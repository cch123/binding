// Copyright 2014 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package binding

import "github.com/cch123/binding/internal/form"

var decoder = form.NewDecoder()

func mapForm(ptr interface{}, form map[string][]string) error {
	//decoder := schema.NewDecoder()
	//decoder.SetAliasTag("form")
	if err := decoder.Decode(ptr, form); err != nil {
		return err
	}

	return nil
}
