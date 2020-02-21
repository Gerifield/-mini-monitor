package docker

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gerifield/mini-monitor/src/checker/config"
)

func TestDockerInitID(t *testing.T) {
	testTable := []struct {
		params map[string]interface{}
		err    error
	}{
		{params: nil, err: nil},
		{params: map[string]interface{}{}, err: nil},
		{params: map[string]interface{}{"testKey1": "testVal1"}, err: nil},
		{params: map[string]interface{}{confID: true}, err: config.ErrLoadFailed},
		{params: map[string]interface{}{confID: "testID"}, err: nil},
	}

	check := New()
	for _, tt := range testTable {
		assert.Equal(t, tt.err, check.Init(tt.params))
	}
}

func TestDockerInitNameRegex(t *testing.T) {
	testTable := []struct {
		params map[string]interface{}
		err    error
	}{
		{params: nil, err: nil},
		{params: map[string]interface{}{}, err: nil},
		{params: map[string]interface{}{"testKey1": "testVal1"}, err: nil},
		{params: map[string]interface{}{confNameRegex: true}, err: config.ErrLoadFailed},
		{params: map[string]interface{}{confNameRegex: "testRegex"}, err: nil},
	}

	check := New()
	for _, tt := range testTable {
		assert.Equal(t, tt.err, check.Init(tt.params))
	}
}

func TestDockerInitImageRegex(t *testing.T) {
	testTable := []struct {
		params map[string]interface{}
		err    error
	}{
		{params: nil, err: nil},
		{params: map[string]interface{}{}, err: nil},
		{params: map[string]interface{}{"testKey1": "testVal1"}, err: nil},
		{params: map[string]interface{}{confImageRegex: true}, err: config.ErrLoadFailed},
		{params: map[string]interface{}{confImageRegex: "testRegex"}, err: nil},
	}

	check := New()
	for _, tt := range testTable {
		assert.Equal(t, tt.err, check.Init(tt.params))
	}
}
