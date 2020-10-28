package hello

const english = "english"
const englishPrefix = "Hello"
const chinese = "chinese"
const chinesePrefix = "你好"

// Hello say hello by language
func Hello(language, name string) string {
	if name == "" {
		name = "world"
	}
	return greetingPrefix(language) + ", " + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case english:
		prefix = englishPrefix
	case chinese:
		prefix = chinesePrefix
	default:
		prefix = englishPrefix
	}
	return
}
