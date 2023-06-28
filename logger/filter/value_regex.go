package filter

import (
	"fmt"
	"regexp"
)

type ValueRegexFilter struct {
	format string
}

func ValueRegex(format string) *ValueRegexFilter {
	return &ValueRegexFilter{format: format}
}

func (f *ValueRegexFilter) DoFilter(value any) any {
	return regexp.MustCompile(f.format).ReplaceAllString(fmt.Sprintf("%v", value), LabelFiltered)
}

func (f *ValueRegexFilter) IsFilter(key string, value any) bool {
	switch value.(type) {
	case string:
		return true
	default:
		return false
	}
}
