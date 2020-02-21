package debug

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {

	testTable := []struct {
		params map[string]interface{}
		err    error
	}{
		{params: nil, err: nil},
		{params: map[string]interface{}{}, err: nil},
		{params: map[string]interface{}{"testKey1": "testVal1"}, err: nil},
		{params: map[string]interface{}{loadFail: "testVal1"}, err: nil},
		{params: map[string]interface{}{loadFail: true}, err: errLoadFailed},
	}

	check := New()
	for _, tt := range testTable {
		assert.Equal(t, tt.err, check.Init(tt.params))
	}
}
