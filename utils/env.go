package utils

import (
	"log"
	"os"
	"strconv"
)

// MustGetEnv : return env variable otherwise logs error message if undefined
func MustGetEnv(env string) string {
	v := os.Getenv(env)
	if v == "" {
		log.Panicf("%s missing", env)
	}
	return v
}

// MustGetEnvBool : return env variable as a bool otherwise log error if undefined
func MustGetEnvBool(env string) bool {
	v := os.Getenv(env)
	if v == "" {
		log.Panicf("%s missing", env)
	}
	b, err := strconv.ParseBool(v)
	if err != nil {
		log.Panicf("Error '%q' while parsing env variable: %s", err, env)
	}
	return b
}
