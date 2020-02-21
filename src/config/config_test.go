package config

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestReadConfig(t *testing.T) {
	f, err := os.Open("test.json")
	require.NoError(t, err)
	defer func() { _ = f.Close() }()

	conf, err := ReadConfig(f)
	assert.NoError(t, err)
	assert.Equal(t, 10*time.Second, conf.CheckTime)

	require.Len(t, conf.Configs, 2)

	assert.Equal(t, CheckConfig{
		Name:   "test1",
		Type:   "debug",
		Config: map[string]interface{}{"testKey": "testVal"},
	}, conf.Configs[0])

	assert.Equal(t, CheckConfig{
		Name:   "test2",
		Type:   "debug",
		Config: nil,
	}, conf.Configs[1])
}
