package filter

import "strings"

type FieldFilter struct {
	targets map[string]string
}

func Field(targets ...string) *FieldFilter {
	fields := make(map[string]string, 0)

	for _, target := range targets {
		fields[target] = target
	}

	return &FieldFilter{targets: fields}
}

func (f *FieldFilter) DoFilter(value any) any {
	return LabelFiltered
}

func (f *FieldFilter) IsFilter(key string, value any) bool {
	return strings.EqualFold(key, f.targets[key])
}
