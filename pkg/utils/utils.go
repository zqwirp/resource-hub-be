package utils

import (
	"fmt"
	"os"
)

func GetEnvWithPanic(key string) string {
	value := os.Getenv(key)
	if value == "" {
		panic(fmt.Sprintf("environment variable %s not set", key))
	}
	return value
}

func GetEnvWithDefaultValueAndPanic(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		value = defaultValue
	}

	if value == "invalid" {
		panic(fmt.Sprintf("critical error: environment variable %s has an invalid value", key))
	}

	return value
}

func GetEnvsWithPanic(args ...string) string {
	for _, key := range args {
		value := os.Getenv(key)
		if value != "" {
			return value
		}
	}
	panic("None of the specified environment variables are set")
}
