package types

import (
	"io/ioutil"
	"net/http"

	"gopkg.in/yaml.v2"
)

type Play struct {
	Hosts SingleOrMulti `yaml:"hosts"`
	Steps []Step        `yaml:"steps"`
}

// Plays the play
func (p *Play) Play(ctx *Context, client *http.Client) {
	for _, step := range p.Steps {
		for _, host := range p.Hosts.Val {
			step.Exec(ctx, client, host)
		}
	}
}

// Creates a new play using a yaml file
func NewPlay(file string) *Play {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	var play Play
	yaml.Unmarshal(bytes, &play)
	return &play
}
