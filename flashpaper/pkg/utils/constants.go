package utils

var SupportedLanguages = map[string]bool{
	"text":       true,
	"go":         true,
	"python":     true,
	"javascript": true,
	"java":       true,
	"html":       true,
	"css":        true,
	"sql":        true,
}

func SanitizeLanguage(lang string) string {
	if SupportedLanguages[lang] {
		return lang
	}
	return "text"
}
