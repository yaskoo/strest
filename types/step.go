package types

import (
	"fmt"
	_ "io/ioutil"
	"net/http"
)

type Step struct {
	Name    string                   `yaml:"name"`
	Url     string                   `yaml:"url"`
	Skip    bool                     `yaml:"skip"`
	Method  string                   `yaml:"method"`
	Headers map[string]SingleOrMulti `yaml:"headers"`
	Body    *Body                    `yaml:"body"`
}

func (s *Step) String() string {
	if s.Name != "" {
		return s.Name
	}
	return fmt.Sprintf("@%s %s", s.Method, s.Url)
}

func (s *Step) Exec(ctx *Context, client *http.Client, host string) {
	if s.Skip {
		fmt.Printf("Step: %s (skipped)\n", s.String())
		return
	}

	req, _ := http.NewRequest(s.Method, host+s.Url, nil)
	for key, value := range s.Headers {
		for _, h := range value.Val {
			req.Header.Add(key, h)
		}
	}

	if s.Body != nil {
		req.Body = s.Body.Reader()
	}

	fmt.Printf("Step: %s\n", s.String())
	client.Do(req)
	// res, _ := client.Do(req)
	// _, _ := ioutil.ReadAll(res.Body)
}
