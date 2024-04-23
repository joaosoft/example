package example

import "os"

func getEnv() string {
	env := os.Getenv("env")
	if env == "" {
		env = "local"
	}

	return env
}
