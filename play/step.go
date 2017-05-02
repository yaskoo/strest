package types

import (
	"bytes"
	"fmt"
	tpl "github.com/yaskoo/strest/template"
	"io/ioutil"
	"net/http"
	"text/template"
)

type Step struct {
	Name     string                   `yaml:"name"`
	Url      string                   `yaml:"url"`
	Skip     bool                     `yaml:"skip"`
	Method   string                   `yaml:"method"`
	Headers  map[string]SingleOrMulti `yaml:"headers"`
	Body     Body                     `yaml:"body"`
	Register []RegVal                 `yaml:"register"`
}

func (s *Step) Exec(ctx *Context, client *http.Client, host string) {
	if s.Skip {
		fmt.Printf("Step (skipped): %s\n", s.Name)
		return
	}

	fmt.Printf("Step: %s\n", s.Name)

	req, _ := http.NewRequest(s.Method, applyTpl(ctx, host+s.Url), nil)
	for key, value := range s.Headers {
		for _, h := range value.Val {
			req.Header.Add(key, applyTpl(ctx, h))
		}
	}

	req.Body = s.Body.Reader()

	// TODO: err handling
	res, _ := client.Do(req)
	b, _ := ioutil.ReadAll(res.Body)

	ctx.Res = &Response{
		Body: string(b),
	}

	for _, reg := range s.Register {
		ctx.Register[reg.Key] = applyTpl(ctx, reg.Val)
	}
}

func applyTpl(ctx *Context, s string) string {
	var t *template.Template
	var err error
	if t, err = template.New("tpl").Funcs(tpl.TplFuncMap).Parse(s); err != nil {
		panic(err)
	}

	var result bytes.Buffer
	if err := t.Execute(&result, ctx); err != nil {
		panic(err)
	}
	return result.String()
}
