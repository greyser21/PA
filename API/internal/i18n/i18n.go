package i18n

import "github.com/gin-gonic/gin"

// Minimal i18n: languages can be extended via resource files without code changes.
var messages = map[string]map[string]string{
	"fr": {"welcome": "Bienvenue sur UpcycleConnect", "error.invalid": "Données invalides"},
	"en": {"welcome": "Welcome to UpcycleConnect", "error.invalid": "Invalid data"},
}

func Middleware(defaultLang string) gin.HandlerFunc {
	return func(c *gin.Context) {
		lang := c.GetHeader("X-Lang")
		if lang == "" {
			lang = c.Query("lang")
		}
		if _, ok := messages[lang]; !ok || lang == "" {
			lang = defaultLang
		}
		c.Set("lang", lang)
		c.Next()
	}
}

func T(lang, key string) string {
	if m, ok := messages[lang]; ok {
		if v, ok := m[key]; ok {
			return v
		}
	}
	if m, ok := messages["fr"]; ok {
		if v, ok := m[key]; ok {
			return v
		}
	}
	return key
}
