package template

import (
	"encoding/base64"
	"text/template"

	"github.com/yaskoo/strest/util/json"
)

var TplFuncMap = template.FuncMap{
	"get_json": json.Get,
	"base64": func(data string) string {
		return base64.StdEncoding.EncodeToString([]byte(data))
	},
}
