package env

import (
	"fmt"
	"os"
)

func GetEnv(env string) (string, error) {
	value := os.Getenv(env)

	if value == "" {
		return "", fmt.Errorf("environment %s was not defined", env)
	}

	return value, nil
}
