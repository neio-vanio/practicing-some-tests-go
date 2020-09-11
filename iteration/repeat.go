package iteration

// Repeat string
func Repeat(text string, n int) (result string) {
	for i := 0; i < n; i++ {
		result += text
	}
	return
}
