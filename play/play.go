package play

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Play struct {
	Hosts SingleOrMulti `yaml:"hosts"`
	Steps []Step        `yaml:"steps"`
}

type Step struct {
	Name     string                   `yaml:"name"`
	Url      string                   `yaml:"url"`
	Skip     bool                     `yaml:"skip"`
	Method   string                   `yaml:"method"`
	Headers  map[string]SingleOrMulti `yaml:"headers"`
	Body     Body                     `yaml:"body"`
	Register []RegVal                 `yaml:"register"`
}

func (p *Play) Load(file string) error {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	yaml.Unmarshal(bytes, p)
	return nil
}
