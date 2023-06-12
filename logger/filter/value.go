package filter

import (
	"fmt"
	"strings"
)

type ValueFilter struct {
	targets map[string]string
}

func Value(targets ...string) *ValueFilter {
	values := make(map[string]string, 0)

	for _, target := range targets {
		values[target] = target
	}

	return &ValueFilter{targets: values}
}

func (f *ValueFilter) DoFilter(value any) any {
	valueKey := fmt.Sprintf("%v", value)

	return strings.ReplaceAll(valueKey, f.targets[valueKey], LabelFiltered)
}

func (f *ValueFilter) IsFilter(key string, value any) bool {
	valueKey := fmt.Sprintf("%v", value)

	return valueKey == f.targets[valueKey]
}
