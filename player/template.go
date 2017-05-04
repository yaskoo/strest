package player

import (
	"bytes"
	"text/template"

	"github.com/yaskoo/strest/play"
	tpl "github.com/yaskoo/strest/template"
)

func Template(ctx *play.Context, s string) (string, error) {
	var t *template.Template
	var err error
	if t, err = template.New("tpl").Funcs(tpl.TplFuncMap).Parse(s); err != nil {
		return "", err
	}

	var result bytes.Buffer
	if err := t.Execute(&result, ctx); err != nil {
		return "", err
	}
	return result.String(), nil
}
