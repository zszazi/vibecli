package env

import (
	"os"
	"strconv"
	"time"
)

func GetString(key, fallback string) string {
	val, ok := os.LookupEnv(key)

	if !ok {
		return fallback
	}

	return val
}

func GetInt(key string, fallback int) int {
	val, ok := os.LookupEnv(key)

	if !ok {
		return fallback
	}

	valAsInt, err := strconv.Atoi(val)
	if err != nil {
		return fallback
	}

	return valAsInt
}

// define string time as per time.Duration
func GetTimeDuration(key, fallback string) time.Duration {

	fallbackVal, err := time.ParseDuration(fallback)
	if err != nil {
		return fallbackVal //TOOD: appropriate handle error
	}

	val, err := time.ParseDuration(key)
	if err != nil {
		return fallbackVal
	}

	return val
}
