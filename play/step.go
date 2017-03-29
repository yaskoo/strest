package play

import (
	"bytes"
	"container/list"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
	"text/template"

	tpl "github.com/yaskoo/strest/template"
)

func (s *Step) Exec(ctx *Context, client *http.Client) {
	req, _ := http.NewRequest(s.Method, s.Url, nil)

	for key, value := range s.Headers {
		for _, h := range value {
			t := template.Must(template.New("header").Funcs(tpl.TplFuncMap).Parse(h))

			var header bytes.Buffer
			t.Execute(&header, ctx)
			req.Header.Add(key, header.String())
		}
	}

	if s.Body != "" {
		req.Body = ioutil.NopCloser(strings.NewReader(s.Body))
	}

	res, _ := client.Do(req)
	resBody, _ := ioutil.ReadAll(res.Body)
	s.Response = &Response{
		Status:  res.StatusCode,
		Headers: map[string][]string(res.Header),
		Body:    string(resBody),
	}

	ctx.Responses[ctx.step] = s.Response
}

func (s *Step) Assert() *list.List {
	diff := list.New()
	if s.Expect == nil {
		return diff
	}

	if s.Response.Status != s.Expect.Status {
		diff.PushBack(DiffEntry{
			Prop:     Status,
			Expected: s.Expect.Status,
			Actual:   s.Response.Status,
		})
	}

	var expected interface{}
	var actual interface{}
	json.Unmarshal([]byte(s.Expect.Body), &expected)
	json.Unmarshal([]byte(s.Response.Body), &actual)
	if !reflect.DeepEqual(expected, actual) {
		diff.PushBack(DiffEntry{
			Prop:     Body,
			Expected: expected,
			Actual:   actual,
		})
	}
	return diff
}

func (s *Step) SetDefaults() {
	if s.Method == "" {
		s.Method = "GET"
	}
}
