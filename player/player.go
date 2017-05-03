package player

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaskoo/strest/play"
)

type Player struct {
	Ctx    *play.Context
	Client *http.Client
}

func New() *Player {
	ctx := &play.Context{
		Register: make(map[string]string),
	}

	return &Player{
		Ctx:    ctx,
		Client: &http.Client{},
	}
}

func (pl *Player) Play(p *play.Play) {
	for _, step := range p.Steps {
		for _, host := range p.Hosts.Val {
			pl.PlayStep(&step, host)
		}
	}
}

func (pl *Player) PlayStep(s *play.Step, host string) {
	if s.Skip {
		fmt.Printf("Step (skipped): %s\n", s.Name)
		return
	}

	fmt.Printf("Step: %s\n", s.Name)

	req, _ := http.NewRequest(s.Method, Template(pl.Ctx, host+s.Url), nil)
	for key, value := range s.Headers {
		for _, h := range value.Val {
			req.Header.Add(key, Template(pl.Ctx, h))
		}
	}

	req.Body = s.Body.Reader()

	// TODO: err handling
	res, _ := pl.Client.Do(req)
	b, _ := ioutil.ReadAll(res.Body)

	pl.Ctx.Res = &play.Response{
		Body: string(b),
	}

	for _, reg := range s.Register {
		pl.Ctx.Register[reg.Key] = Template(pl.Ctx, reg.Val)
	}
}
