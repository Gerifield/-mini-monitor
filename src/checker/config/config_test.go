package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReadConfig(t *testing.T) {
	f, err := os.Open("test.json")
	require.NoError(t, err)
	defer func() { _ = f.Close() }()

	checks, err := ReadConfig(f)
	assert.NoError(t, err)
	require.Len(t, checks.Configs, 2)

	assert.Equal(t, CheckConfig{
		Name:   "test1",
		Type:   "debug",
		Config: map[string]interface{}{"testKey": "testVal"},
	}, checks.Configs[0])

	assert.Equal(t, CheckConfig{
		Name:   "test2",
		Type:   "debug",
		Config: nil,
	}, checks.Configs[1])
}
