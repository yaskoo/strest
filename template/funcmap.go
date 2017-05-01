package template

import (
	"text/template"

	"github.com/yaskoo/strest/util/json"
)

var TplFuncMap = template.FuncMap{
	"get_json": jsonq.Get,
}
