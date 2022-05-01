package subseq

// MaxSubstring возвращаает максимальную общую подстроку
func MaxSubstring(s1, s2 string) string {
	return string(MaxSubset([]rune(s1), []rune(s2)))
}

// MaxSubsequence возвращает максимальную общую подпоследовательность символов
func MaxSubsequence(s1, s2 string) string {
	return string(MaxSubsetSequence([]rune(s1), []rune(s2)))
}
