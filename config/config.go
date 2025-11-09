package config

import "errors"

// Config defines a generic configuration interface.
type Config interface {
	// Get returns a value by key.
	Get(key string) (interface{}, error)
	// GetString returns a string value by key.
	GetString(key string) (string, error)
	// GetInt returns an int value by key.
	GetInt(key string) (int, error)
	// GetBool returns a bool value by key.
	GetBool(key string) (bool, error)
	// Unmarshal populates the given struct with the configuration values.
	Unmarshal(out interface{}) error
}

var ErrKeyNotFound = errors.New("key not found")
