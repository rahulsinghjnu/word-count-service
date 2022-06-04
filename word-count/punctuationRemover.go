package word_count

import (
	"fmt"
	"regexp"
)

type PunctuationRemoverInterface interface {
	RemovePunctuation(string) (string, error)
}

type RegexPunctuationRemover struct {
}

func (regPunctRemover RegexPunctuationRemover) RemovePunctuation(str string) (string, error) {
	var punctuationPattern *regexp.Regexp
	punctuationPattern, err := regexp.Compile("[!,.]")
	if err != nil {
		fmt.Println("Failed to compile !!")
		return "", err
	}
	respStr := punctuationPattern.ReplaceAllString(str, "")
	return respStr, nil
}
