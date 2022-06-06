package word_count

import (
	"fmt"
	"regexp"
)

type PunctuationFilter struct {
}

func (punctFilter PunctuationFilter) Filter(str string) (string, error) {
	var punctuationPattern *regexp.Regexp
	punctuationPattern, err := regexp.Compile("[!,.?]")
	if err != nil {
		fmt.Println("Failed to compile !!")
		return "", err
	}
	respStr := punctuationPattern.ReplaceAllString(str, "")
	return respStr, nil
}

func NewPunctFilter() WordFilter {
	return PunctuationFilter{}
}
