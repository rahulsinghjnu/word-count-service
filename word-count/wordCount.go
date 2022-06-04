package word_count

import (
	"strings"
)

type PunctuationRemovedWordCountService struct {
	PunctuationRemover PunctuationRemoverInterface
}

func (wordCountService PunctuationRemovedWordCountService) CountWords(str string) (map[string]int, error) {
	punctuationFreeStr, err := wordCountService.PunctuationRemover.RemovePunctuation(str)
	if err != nil {
		return map[string]int{}, err
	}
	words := strings.Fields(punctuationFreeStr)
	wordCounts := make(map[string]int)
	for _, word := range words {
		wordCounts[strings.ToUpper(word)] += 1
	}
	return wordCounts, nil
}

func New() WordCountInterface {
	return PunctuationRemovedWordCountService{PunctuationRemover: RegexPunctuationRemover{}}
}
