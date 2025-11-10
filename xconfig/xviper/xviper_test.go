package xviper_test

import (
	"os"
	"testing"

	"github.com/agussyahrilmubarok/gohelp/xconfig"
	"github.com/agussyahrilmubarok/gohelp/xconfig/xviper"
	"github.com/stretchr/testify/assert"
)

func createTempConfigFile(t *testing.T, filename, content string) string {
	t.Helper()
	tmpFile, err := os.CreateTemp("", filename)
	assert.NoError(t, err)

	_, err = tmpFile.WriteString(content)
	assert.NoError(t, err)

	err = tmpFile.Close()
	assert.NoError(t, err)

	return tmpFile.Name()
}

func TestViperConfig(t *testing.T) {
	yamlContent := `
app:
  name: "test-app"
  port: 8080
  debug: true
database:
  host: "localhost"
  port: 5432
`

	configFile := createTempConfigFile(t, "config-*.yaml", yamlContent)
	defer os.Remove(configFile)

	v, err := xviper.NewConfig(configFile)
	assert.NoError(t, err)
	assert.NotNil(t, v)

	appName, err := v.GetString("app.name")
	assert.NoError(t, err)
	assert.Equal(t, "test-app", appName)

	appPort, err := v.GetInt("app.port")
	assert.NoError(t, err)
	assert.Equal(t, 8080, appPort)

	debug, err := v.GetBool("app.debug")
	assert.NoError(t, err)
	assert.True(t, debug)

	val, err := v.Get("database.host")
	assert.NoError(t, err)
	assert.Equal(t, "localhost", val)

	_, err = v.GetString("app.nonexistent")
	assert.ErrorIs(t, err, xconfig.ErrKeyNotFound)

	type Config struct {
		App struct {
			Name  string `mapstructure:"name"`
			Port  int    `mapstructure:"port"`
			Debug bool   `mapstructure:"debug"`
		} `mapstructure:"app"`

		Database struct {
			Host string `mapstructure:"host"`
			Port int    `mapstructure:"port"`
		} `mapstructure:"database"`
	}

	var ac Config
	err = v.Unmarshal(&ac)
	assert.NoError(t, err)
	assert.Equal(t, "test-app", ac.App.Name)
	assert.Equal(t, 8080, ac.App.Port)
	assert.True(t, ac.App.Debug)
}
