package config

type AppEnvironment string

const (
	appEnvironmentDevelopment = "development"
	appEnvironmentProduction  = "production"
)

func IsProduction(env AppEnvironment) bool {
	return env == appEnvironmentProduction
}
