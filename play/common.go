package types

type RegVal struct {
	Key string `yaml:"key"`
	Val string `yaml:"val"`
}

type SingleOrMulti struct {
	Val []string
}

func (sm *SingleOrMulti) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var multi []string
	err := unmarshal(&multi)
	if err != nil {
		var single string
		err := unmarshal(&single)
		if err != nil {
			return err
		}
		sm.Val = make([]string, 1)
		sm.Val[0] = single
	} else {
		sm.Val = multi
	}
	return nil
}
