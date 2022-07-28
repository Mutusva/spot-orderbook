package env

import "os"

type RedisConfig struct {
	Host     string
	Password string
	Port     string
}

func GetRedisConfig(env string) *RedisConfig {
	setEnvironmentVariable(env)
	return &RedisConfig{
		Host:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PWD"),
		Port:     os.Getenv("REDIS_PORT"),
	}
}

func (r *RedisConfig) IsProd() {

}

func setEnvironmentVariable(env string) {
	if env == "prod" {
		os.Setenv("REDIS_HOST", "172.30.29.69")
		os.Setenv("REDIS_PWD", "Hjyloa4n9l")
		os.Setenv("REDIS_PORT", "6379")
	} else {
		os.Setenv("REDIS_HOST", "localhost")
		os.Setenv("REDIS_PWD", "")
		os.Setenv("REDIS_PORT", "6379")
	}
}
