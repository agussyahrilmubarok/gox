package envconfig

import (
	"errors"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/agussyahrilmubarok/gohelp/config"
)

// EnvConfig provides access to configuration values from environment variables.
type EnvConfig struct{}

// New creates a new EnvConfig instance.
func New() *EnvConfig {
	return &EnvConfig{}
}

// Get retrieves the value of the given environment variable as interface{}.
// Returns ErrKeyNotFound if the key does not exist.
func (e *EnvConfig) Get(key string) (interface{}, error) {
	val, ok := os.LookupEnv(key)
	if !ok {
		return nil, config.ErrKeyNotFound
	}
	return val, nil
}

// GetString retrieves the value of the given environment variable as string.
// Returns ErrKeyNotFound if the key does not exist.
func (e *EnvConfig) GetString(key string) (string, error) {
	val, ok := os.LookupEnv(key)
	if !ok {
		return "", config.ErrKeyNotFound
	}
	return val, nil
}

// GetInt retrieves the value of the given environment variable as int.
// Returns ErrKeyNotFound if the key does not exist or conversion fails.
func (e *EnvConfig) GetInt(key string) (int, error) {
	valStr, ok := os.LookupEnv(key)
	if !ok {
		return 0, config.ErrKeyNotFound
	}

	val, err := strconv.Atoi(valStr)
	if err != nil {
		return 0, err
	}
	return val, nil
}

// GetBool retrieves the value of the given environment variable as bool.
// Returns ErrKeyNotFound if the key does not exist or conversion fails.
func (e *EnvConfig) GetBool(key string) (bool, error) {
	valStr, ok := os.LookupEnv(key)
	if !ok {
		return false, config.ErrKeyNotFound
	}

	val, err := strconv.ParseBool(valStr)
	if err != nil {
		return false, err
	}
	return val, nil
}

// Unmarshal maps environment variables into the fields of a struct.
// Fields can use `mapstructure` tags to specify the corresponding environment variable name.
// Only string, int, and bool types are supported.
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
		mapTag := field.Tag.Get("mapstructure")
		if mapTag == "" {
			mapTag = strings.ToUpper(field.Name)
		}

		envVal, ok := os.LookupEnv(mapTag)
		if !ok {
			continue
		}

		fieldVal := v.Field(i)
		if !fieldVal.CanSet() {
			continue
		}

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
		// Extension point: handle float, slices, nested structs, etc.
		default:
			continue
		}
	}

	return nil
}
