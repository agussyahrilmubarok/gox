package viperconfig

import (
	"path/filepath"
	"strings"

	"github.com/agussyahrilmubarok/gohelp/config"
	"github.com/spf13/viper"
)

// ViperConfig provides access to configuration values using Viper.
type ViperConfig struct {
	v *viper.Viper
}

// New creates a new ViperConfig instance with the specified config file.
// Supports YAML, JSON, or defaults to YAML.
func New(configFile string) (*ViperConfig, error) {
	v := viper.New()
	v.SetConfigFile(configFile)

	ext := strings.ToLower(filepath.Ext(configFile))
	switch ext {
	case ".yaml", ".yml":
		v.SetConfigType("yaml")
	case ".json":
		v.SetConfigType("json")
	default:
		// Default to YAML if file extension is unrecognized
		v.SetConfigType("yaml")
	}

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	return &ViperConfig{v: v}, nil
}

// Get retrieves the value of the given key as interface{}.
// Returns ErrKeyNotFound if the key does not exist.
func (c *ViperConfig) Get(key string) (interface{}, error) {
	if !c.v.IsSet(key) {
		return nil, config.ErrKeyNotFound
	}
	return c.v.Get(key), nil
}

// GetString retrieves the value of the given key as string.
// Returns ErrKeyNotFound if the key does not exist.
func (c *ViperConfig) GetString(key string) (string, error) {
	if !c.v.IsSet(key) {
		return "", config.ErrKeyNotFound
	}
	return c.v.GetString(key), nil
}

// GetInt retrieves the value of the given key as int.
// Returns ErrKeyNotFound if the key does not exist.
func (c *ViperConfig) GetInt(key string) (int, error) {
	if !c.v.IsSet(key) {
		return 0, config.ErrKeyNotFound
	}
	return c.v.GetInt(key), nil
}

// GetBool retrieves the value of the given key as bool.
// Returns ErrKeyNotFound if the key does not exist.
func (c *ViperConfig) GetBool(key string) (bool, error) {
	if !c.v.IsSet(key) {
		return false, config.ErrKeyNotFound
	}
	return c.v.GetBool(key), nil
}

// Unmarshal populates a struct with configuration values from the loaded config file.
// The struct fields must match config keys or use `mapstructure` tags.
func (c *ViperConfig) Unmarshal(out interface{}) error {
	return c.v.Unmarshal(out)
}
