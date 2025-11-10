package xenv_test

import (
	"os"
	"testing"

	"github.com/agussyahrilmubarok/gohelp/xconfig"
	"github.com/agussyahrilmubarok/gohelp/xconfig/xenv"
	"github.com/stretchr/testify/assert"
)

func TestEnvConfig(t *testing.T) {
	os.Setenv("APP_NAME", "test-app")
	os.Setenv("APP_PORT", "8080")
	os.Setenv("DEBUG_MODE", "true")
	defer func() {
		os.Unsetenv("APP_NAME")
		os.Unsetenv("APP_PORT")
		os.Unsetenv("DEBUG_MODE")
	}()

	e := xenv.NewConfig()
	assert.NotNil(t, e)

	val, err := e.Get("APP_NAME")
	assert.NoError(t, err)
	assert.Equal(t, "test-app", val)

	_, err = e.Get("NON_EXISTENT")
	assert.ErrorIs(t, err, xconfig.ErrKeyNotFound)

	strVal, err := e.GetString("APP_NAME")
	assert.NoError(t, err)
	assert.Equal(t, "test-app", strVal)

	_, err = e.GetString("NON_EXISTENT")
	assert.ErrorIs(t, err, xconfig.ErrKeyNotFound)

	intVal, err := e.GetInt("APP_PORT")
	assert.NoError(t, err)
	assert.Equal(t, 8080, intVal)

	os.Setenv("APP_PORT", "not-an-int")
	_, err = e.GetInt("APP_PORT")
	assert.Error(t, err)
	os.Setenv("APP_PORT", "8080")

	boolVal, err := e.GetBool("DEBUG_MODE")
	assert.NoError(t, err)
	assert.True(t, boolVal)

	os.Setenv("DEBUG_MODE", "not-bool")
	_, err = e.GetBool("DEBUG_MODE")
	assert.Error(t, err)
	os.Setenv("DEBUG_MODE", "true")

	type Config struct {
		AppName   string `mapstructure:"APP_NAME"`
		AppPort   int    `mapstructure:"APP_PORT"`
		DebugMode bool   `mapstructure:"DEBUG_MODE"`
	}

	var cfg Config
	err = e.Unmarshal(&cfg)
	assert.NoError(t, err)
	assert.Equal(t, "test-app", cfg.AppName)
	assert.Equal(t, 8080, cfg.AppPort)
	assert.True(t, cfg.DebugMode)

	err = e.Unmarshal(cfg)
	assert.Error(t, err)

	var x int
	err = e.Unmarshal(&x)
	assert.Error(t, err)
}
