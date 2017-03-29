package play

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"gopkg.in/yaml.v2"
)

// Plays the play
func (p *Play) Exec(ctx *Context, client *http.Client) {
	ok := true
	for i, step := range p.Steps {
		ctx.step = i
		step.SetDefaults()
		step.Exec(ctx, client)

		if p.testMode {
			diff := step.Assert()
			ok = ok && diff.Len() == 0

			for e := diff.Front(); e != nil; e = e.Next() {
				fmt.Println(e.Value)
			}
		}
	}

	if p.testMode && !ok {
		os.Exit(1)
	}
}

// Plays the play in test mode i.e assert expectations
func (p *Play) ExecTest(ctx *Context, client *http.Client) {
	p.testMode = true
	p.Exec(ctx, client)
}

func New(file string) *Play {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	var play Play
	yaml.Unmarshal(bytes, &play)
	return &play
}
