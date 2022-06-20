package utils

import "github.com/tkanos/gonfig"

func LoadConfigurations[T any](path string) (*T, error) {
	var c T

	err := gonfig.GetConf(path, c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
