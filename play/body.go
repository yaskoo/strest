package play

import (
	"io"
	"io/ioutil"
	"net/url"
	"strings"
)

type BodyType int

const (
	Text BodyType = iota
	Form
)

type Body struct {
	data string
	Type BodyType `yaml:"-"`
}

func (r *Body) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var v string
	if err := unmarshal(&v); err == nil {
		r.data = v
		r.Type = Text
		return nil
	}

	var formBody struct {
		Form map[string]SingleOrMulti `yaml:"form"`
	}

	if err := unmarshal(&formBody); err == nil {
		data := url.Values{}
		for k, v := range formBody.Form {
			for _, i := range v.Val {
				data.Add(k, i)
			}
		}

		r.data = data.Encode()
		r.Type = Form
		return nil
	}
	return nil
}

func (b *Body) Reader() io.ReadCloser {
	return ioutil.NopCloser(strings.NewReader(b.data))
}

func (b *Body) String() string {
	return b.data
}
