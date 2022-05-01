package levenshtein

// String возвращает последовательность правок для превращения строки s1 в строку s2
func String(s1, s2 string) []Edit {
	return Levenshtein([]rune(s1), []rune(s2), nil)
}

// Strings возвращает последовательность правок для превращения слайса строк s1 в слайс строк s2
func Strings(s1, s2 []string) []Edit {
	return Levenshtein(s1, s2, nil)
}
