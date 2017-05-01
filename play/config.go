package types

import (
	"os"
)

var Conf Configuration

type Configuration struct {
	file   string
	Output string
}

func init() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	Conf.Output = wd
}
