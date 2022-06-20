package utils

import "github.com/tkanos/gonfig"

type Configuration interface {
}

func LoadConfigurations[T Configuration](path string) (Configuration, error) {
	var c T

	err := gonfig.GetConf(path, c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
