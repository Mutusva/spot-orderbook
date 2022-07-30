package env

import "os"

type EnvType string

var (
	PROD EnvType = "prod"
	DEV  EnvType = "dev"
)

type AppConfig struct {
	RedisHost     string
	RedisPassword string
	RedisPort     string
}

// GetRedisConfig TODO: read the config from file
func GetRedisConfig(env string) *AppConfig {
	setEnvironmentVariable(env)
	return &AppConfig{
		RedisHost:     os.Getenv("REDIS_HOST"),
		RedisPassword: os.Getenv("REDIS_PWD"),
		RedisPort:     os.Getenv("REDIS_PORT"),
	}
}

func setEnvironmentVariable(env string) {
	if env == string(PROD) {
		os.Setenv("REDIS_HOST", "172.30.29.69")
		os.Setenv("REDIS_PWD", "Hjyloa4n9lkM4aelooVie8lai2euBhodse")
		os.Setenv("REDIS_PORT", "6379")
	} else {
		os.Setenv("REDIS_HOST", "localhost")
		os.Setenv("REDIS_PWD", "")
		os.Setenv("REDIS_PORT", "6379")
	}
	os.Setenv("ENV", env)
}

func IsProd() bool {
	return os.Getenv("ENV") == string(PROD)
}
