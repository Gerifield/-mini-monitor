package loader

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gerifield/mini-monitor/src/config"
)

var (
	testConfig = map[string]interface{}{
		"key1": "val1",
		"key2": 2.0,
		"key3": true,
		"key4": 2,
	}
)

func TestConfigString(t *testing.T) {
	testTable := []struct {
		key      string
		response string
		err      error
	}{
		{"keyNone", "", nil},
		{"key1", "val1", nil},
		{"key2", "", config.ErrLoadFailed},
	}

	for _, tt := range testTable {
		resp, err := ConfigString(testConfig, tt.key)
		assert.Equal(t, tt.response, resp)
		assert.Equal(t, tt.err, err)
	}
}

func TestConfigBool(t *testing.T) {
	testTable := []struct {
		key      string
		response bool
		err      error
	}{
		{"keyNone", false, nil},
		{"key3", true, nil},
		{"key1", false, config.ErrLoadFailed},
	}

	for _, tt := range testTable {
		resp, err := ConfigBool(testConfig, tt.key)
		assert.Equal(t, tt.response, resp)
		assert.Equal(t, tt.err, err)
	}
}

func TestConfigInt(t *testing.T) {
	testTable := []struct {
		key      string
		response int
		err      error
	}{
		{"keyNone", 0, nil},
		{"key2", 2, nil},
		{"key1", 0, config.ErrLoadFailed},
		{"key4", 2, nil},
	}

	for _, tt := range testTable {
		resp, err := ConfigInt(testConfig, tt.key)
		assert.Equal(t, tt.response, resp)
		assert.Equal(t, tt.err, err)
	}
}
