package config

import "os"

type Config struct {
	Addr           string
	GinMode        string
	StripeKey      string
	OneSignalAppID string
	OneSignalKey   string
	DefaultLang    string
}

func Load() Config {
	return Config{
		Addr:           getenv("API_ADDR", ":8080"),
		GinMode:        getenv("GIN_MODE", "debug"),
		StripeKey:      os.Getenv("STRIPE_KEY"),
		OneSignalAppID: os.Getenv("ONESIGNAL_APP_ID"),
		OneSignalKey:   os.Getenv("ONESIGNAL_API_KEY"),
		DefaultLang:    getenv("DEFAULT_LANG", "fr"),
	}
}

func getenv(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}
