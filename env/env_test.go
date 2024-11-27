package env

import (
	"os"
	"testing"
	"time"
)

func TestGetEnv(t *testing.T) {
	{
		// string
		os.Setenv("STRING", "test")
		strVal := GetEnv("STRING", "default")
		if strVal != "test" {
			t.Errorf("Expected 'test', got %s", strVal)
		}
		strVal = GetEnv("STRING_NONE", "default")
		if strVal != "default" {
			t.Errorf("Expected 'default', got %s", strVal)
		}
	}

	{
		// int
		os.Setenv("INT", "123")
		intVal := GetEnv("INT", 1)
		if intVal != 123 {
			t.Errorf("Expected 123, got %d", intVal)
		}
		intVal = GetEnv("INT_NONE", 1)
		if intVal != 1 {
			t.Errorf("Expected 1, got %d", intVal)
		}
	}

	{
		// bool
		os.Setenv("BOOL", "true")
		boolVal := GetEnv("BOOL", false)
		if boolVal != true {
			t.Errorf("Expected true, got %t", boolVal)
		}
		boolVal = GetEnv("BOOL_NONE", false)
		if boolVal != false {
			t.Errorf("Expected false, got %t", boolVal)
		}
	}

	{
		// float
		os.Setenv("FLOAT", "3.14")
		floatVal := GetEnv("FLOAT", 1.23)
		if floatVal != 3.14 {
			t.Errorf("Expected 3.14, got %f", floatVal)
		}
		floatVal = GetEnv("FLOAT_NONE", 1.23)
		if floatVal != 1.23 {
			t.Errorf("Expected 1.23, got %f", floatVal)
		}
	}

	{
		// duration
		os.Setenv("DURATION", "1h")
		durationVal := GetEnv("DURATION", time.Second)
		if durationVal != time.Hour {
			t.Errorf("Expected 1h, got %s", durationVal)
		}
		durationVal = GetEnv("DURATION_NONE", time.Second)
		if durationVal != time.Second {
			t.Errorf("Expected 1s, got %s", durationVal)
		}
	}
}
