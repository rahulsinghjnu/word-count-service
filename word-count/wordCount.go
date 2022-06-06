package word_count

import (
	"strings"
)

// Word Count Service which implements the Word Count interface.
// It includes list of filters which pre-process the given text
type WordCountService struct {
	Filter []WordFilter
}

func (wordCountService WordCountService) CountWords(str string) (map[string]int, error) {
	filteredString := str
	var err error
	// Apply the filters on requested text.
	for _, wordFilter := range wordCountService.Filter {
		filteredString, err = wordFilter.Filter(filteredString)
		if err != nil {
			return map[string]int{}, err
		}
	}
	// Split the text into words.
	words := strings.Fields(filteredString)

	wordCounts := make(map[string]int)
	for _, word := range words {
		wordCounts[strings.ToUpper(word)] += 1
	}
	return wordCounts, nil
}

func New() WordCountInterface {
	filters := make([]WordFilter, 0)
	return WordCountService{Filter: filters}
}

func (wcs WordCountService) Add(wordFilter WordFilter) WordCountService {
	wcs.Filter = append(wcs.Filter, wordFilter)
	return wcs
}
