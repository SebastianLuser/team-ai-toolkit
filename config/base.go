package config

// BaseConfig contains configuration fields shared across all Educabot backend services.
// Each project embeds this and adds its own fields.
type BaseConfig struct {
	Port           string
	Env            string // local, develop, staging, prod
	DatabaseURL    string
	AuthPublicKey  string
	AllowedOrigins []string
	BugsnagAPIKey  string
}

// LoadBase reads the common environment variables into a BaseConfig.
func LoadBase() BaseConfig {
	return BaseConfig{
		Port:           EnvOr("PORT", "8080"),
		Env:            EnvOr("ENV", "local"),
		DatabaseURL:    MustEnv("DATABASE_URL"),
		AuthPublicKey:  MustEnv("AUTH_PUBLIC_KEY"),
		AllowedOrigins: EnvSplit("ALLOWED_ORIGINS", ",", []string{"*"}),
		BugsnagAPIKey:  EnvOr("API_KEY_BUGSNAG", ""),
	}
}

// IsProduction returns true if the environment is prod or staging.
func (c BaseConfig) IsProduction() bool {
	return c.Env == "prod" || c.Env == "staging"
}
