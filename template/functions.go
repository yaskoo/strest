package template

import (
	"text/template"

	"github.com/yaskoo/strest/util/jsonq"
)

var TplFuncMap = template.FuncMap{
	"jsonq": jsonq.Get,
}
