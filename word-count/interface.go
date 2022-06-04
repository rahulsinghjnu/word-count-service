package word_count

type WordCountInterface interface {
	CountWords(string) (map[string]int, error)
}
