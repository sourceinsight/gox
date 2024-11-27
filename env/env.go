package env

import (
	"os"
	"strconv"
	"time"
)

// EnvType is a type that can be used as an environment variable.
type EnvType interface {
	~string | ~int | ~bool | ~float64 | time.Duration
}

// GetEnv returns the value of an environment variable if it exists, otherwise it returns the default value.
func GetEnv[T EnvType](key string, defaultValue T) T {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	switch any(defaultValue).(type) {
	case string:
		return any(value).(T)
	case int:
		intVal, err := strconv.Atoi(value)
		if err != nil || value == "" {
			return defaultValue
		}
		return any(intVal).(T)
	case bool:
		boolVal, err := strconv.ParseBool(value)
		if err != nil || value == "" {
			return defaultValue
		}
		return any(boolVal).(T)
	case float64:
		floatVal, err := strconv.ParseFloat(value, 64)
		if err != nil || value == "" {
			return defaultValue
		}
		return any(floatVal).(T)
	case time.Duration:
		durationVal, err := time.ParseDuration(value)
		if err != nil || value == "" {
			return defaultValue
		}
		return any(durationVal).(T)
	}
	return defaultValue
}
