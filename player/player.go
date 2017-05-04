package player

import (
	"fmt"
	"io/ioutil"

	"github.com/yaskoo/strest/play"
)

type Player struct {
	Ctx    *play.Context
	Client HttpClient
}

func New() *Player {
	ctx := &play.Context{
		Register: make(map[string]string),
	}

	return &Player{
		Ctx:    ctx,
		Client: Client(),
	}
}

func (pl *Player) Play(p *play.Play) {
	for _, step := range p.Steps {
		fmt.Printf("Step: %s\n", step.Name)
		if step.Skip {
			fmt.Println("[skipped]")
			continue
		}

		for _, host := range p.Hosts.Val {
			if err := pl.PlayStep(&step, host); err == nil {
				fmt.Println("[ok]")
			} else {
				fmt.Printf("[error: %s]\n", err.Error())
			}
		}
	}
}

func (pl *Player) PlayStep(s *play.Step, host string) error {
	req, err := pl.Client.Request(pl.Ctx, s, host)
	if err != nil {
		return err
	}

	res, _ := pl.Client.Do(req)
	b, _ := ioutil.ReadAll(res.Body)

	pl.Ctx.Res = &play.Response{
		Body: string(b),
	}

	for _, reg := range s.Register {
		str, err := Template(pl.Ctx, reg.Val)
		if err != nil {
			return err
		}
		pl.Ctx.Register[reg.Key] = str
	}
	return nil
}
