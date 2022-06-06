package word_count

type WordCountInterface interface {
	CountWords(string) (map[string]int, error)
}

type WordFilter interface {
	Filter(string) (string, error)
}
