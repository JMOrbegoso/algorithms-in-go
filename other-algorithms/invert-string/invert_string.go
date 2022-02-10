package invert_string

func invert(originalText string) string {
	inverse := ""
	length := len(originalText)

	runes := []rune(originalText)

	for i := (length - 1); i >= 0; i-- {
		inverse += string(runes[i])
	}

	return inverse
}
