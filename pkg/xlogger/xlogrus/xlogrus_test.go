package xlogrus_test

import (
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/agussyahrilmubarok/gox/pkg/xlogger/xlogrus"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestNewLogger_FileAndStdout(t *testing.T) {
	tmpDir, err := ioutil.TempDir("", "logtest")
	assert.NoError(t, err)
	defer os.RemoveAll(tmpDir)

	logPath := filepath.Join(tmpDir, "test.log")

	logger, err := xlogrus.NewLogger(logPath, "info")
	assert.NoError(t, err)
	assert.NotNil(t, logger)

	logger.Info("hello world")

	content, err := ioutil.ReadFile(logPath)
	assert.NoError(t, err)
	assert.True(t, strings.Contains(string(content), "hello world"))
}

func TestNewLogger_InvalidLevel_DefaultsToInfo(t *testing.T) {
	tmpDir, err := ioutil.TempDir("", "logtest")
	assert.NoError(t, err)
	defer os.RemoveAll(tmpDir)

	logPath := filepath.Join(tmpDir, "test.log")

	logger, err := xlogrus.NewLogger(logPath, "invalid-level")
	assert.NoError(t, err)
	assert.NotNil(t, logger)
	assert.Equal(t, logrus.InfoLevel, logger.GetLevel())
}

func TestNewLogger_CreatesDirectory(t *testing.T) {
	tmpDir, err := ioutil.TempDir("", "logtest")
	assert.NoError(t, err)
	defer os.RemoveAll(tmpDir)

	nestedDir := filepath.Join(tmpDir, "nested", "dir")
	logPath := filepath.Join(nestedDir, "test.log")

	_, err = xlogrus.NewLogger(logPath, "info")
	assert.NoError(t, err)

	info, err := os.Stat(nestedDir)
	assert.NoError(t, err)
	assert.True(t, info.IsDir())
}

func TestNewLogger_CaptureStdout(t *testing.T) {
	tmpDir, err := ioutil.TempDir("", "logtest")
	assert.NoError(t, err)
	defer os.RemoveAll(tmpDir)

	logPath := filepath.Join(tmpDir, "test.log")

	// Capture stdout
	var buf bytes.Buffer
	stdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	logger, err := xlogrus.NewLogger(logPath, "info")
	assert.NoError(t, err)

	logger.Info("capture stdout")

	w.Close()
	os.Stdout = stdout

	_, _ = buf.ReadFrom(r)
	assert.True(t, strings.Contains(buf.String(), "capture stdout"))
}
