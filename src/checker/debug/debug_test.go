package debug

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gerifield/mini-monitor/src/config"
)

func TestInitLoadFail(t *testing.T) {
	testTable := []struct {
		params map[string]interface{}
		err    error
	}{
		{params: nil, err: nil},
		{params: map[string]interface{}{}, err: nil},
		{params: map[string]interface{}{"testKey1": "testVal1"}, err: nil},
		{params: map[string]interface{}{loadFail: "testVal1"}, err: nil},
		{params: map[string]interface{}{loadFail: true}, err: config.ErrLoadFailed},
	}

	check := New()
	for _, tt := range testTable {
		assert.Equal(t, tt.err, check.Init(tt.params))
	}
}

func TestInitCheckFail(t *testing.T) {
	testTable := []struct {
		params map[string]interface{}
		err    error
	}{
		{params: nil, err: nil},
		{params: map[string]interface{}{}, err: nil},
		{params: map[string]interface{}{"testKey1": "testVal1"}, err: nil},
		{params: map[string]interface{}{checkFail: true}, err: nil},
		{params: map[string]interface{}{checkFail: "testVal1"}, err: config.ErrLoadFailed},
	}

	check := New()
	for _, tt := range testTable {
		assert.Equal(t, tt.err, check.Init(tt.params))
	}
}
