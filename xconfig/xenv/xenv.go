package xenv

import (
	"errors"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/agussyahrilmubarok/gohelp/xconfig"
)

// EnvConfig provides access to environment variables as configuration values.
type EnvConfig struct{}

// NewConfig creates a new instance of EnvConfig.
func NewConfig() *EnvConfig {
	return &EnvConfig{}
}

// Get returns the value of the given environment variable as interface{}.
// Returns ErrKeyNotFound if the key does not exist.
func (e *EnvConfig) Get(key string) (interface{}, error) {
	val, ok := os.LookupEnv(key)
	if !ok {
		return nil, xconfig.ErrKeyNotFound
	}
	return val, nil
}

// GetString returns the value of the given environment variable as a string.
// Returns ErrKeyNotFound if the key does not exist.
func (e *EnvConfig) GetString(key string) (string, error) {
	val, ok := os.LookupEnv(key)
	if !ok {
		return "", xconfig.ErrKeyNotFound
	}
	return val, nil
}

// GetInt returns the value of the given environment variable as an int.
// Returns ErrKeyNotFound if the key does not exist or conversion fails.
func (e *EnvConfig) GetInt(key string) (int, error) {
	valStr, ok := os.LookupEnv(key)
	if !ok {
		return 0, xconfig.ErrKeyNotFound
	}

	val, err := strconv.Atoi(valStr)
	if err != nil {
		return 0, err
	}
	return val, nil
}

// GetBool returns the value of the given environment variable as a bool.
// Returns ErrKeyNotFound if the key does not exist or conversion fails.
func (e *EnvConfig) GetBool(key string) (bool, error) {
	valStr, ok := os.LookupEnv(key)
	if !ok {
		return false, xconfig.ErrKeyNotFound
	}

	val, err := strconv.ParseBool(valStr)
	if err != nil {
		return false, err
	}
	return val, nil
}

// Unmarshal maps environment variables to the fields of a struct.
// Fields can use `mapstructure` tags to specify the environment variable name.
// Only string, int, and bool types are supported currently.
func (e *EnvConfig) Unmarshal(out interface{}) error {
	v := reflect.ValueOf(out)
	if v.Kind() != reflect.Ptr || v.IsNil() {
		return errors.New("out must be a non-nil pointer to a struct")
	}

	v = v.Elem()
	if v.Kind() != reflect.Struct {
		return errors.New("out must be a pointer to a struct")
	}

	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		// Use mapstructure tag or uppercase field name as env variable key
		mapTag := field.Tag.Get("mapstructure")
		if mapTag == "" {
			mapTag = strings.ToUpper(field.Name)
		}

		// Lookup environment variable
		envVal, ok := os.LookupEnv(mapTag)
		if !ok {
			continue
		}

		fieldVal := v.Field(i)
		if !fieldVal.CanSet() {
			continue
		}

		// Set field based on its kind
		switch fieldVal.Kind() {
		case reflect.String:
			fieldVal.SetString(envVal)
		case reflect.Int, reflect.Int64, reflect.Int32:
			if intVal, err := strconv.Atoi(envVal); err == nil {
				fieldVal.SetInt(int64(intVal))
			}
		case reflect.Bool:
			if boolVal, err := strconv.ParseBool(envVal); err == nil {
				fieldVal.SetBool(boolVal)
			}
		// Extension point: add support for float, slices, nested structs, etc.
		default:
			continue
		}
	}

	return nil
}
