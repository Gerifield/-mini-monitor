package http

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/gerifield/mini-monitor/src/config"
)

func TestDockerInitParams(t *testing.T) {
	testTable := []struct {
		params map[string]interface{}
		err    error
	}{
		{params: nil, err: nil},
		{params: map[string]interface{}{}, err: nil},
		{params: map[string]interface{}{"testKey1": "testVal1"}, err: nil},

		// Method
		{params: map[string]interface{}{confMethod: true}, err: config.ErrLoadFailed},
		{params: map[string]interface{}{confMethod: "testMethod"}, err: nil},

		// URL
		{params: map[string]interface{}{confURL: true}, err: config.ErrLoadFailed},
		{params: map[string]interface{}{confURL: "testURL"}, err: nil},

		// Body
		{params: map[string]interface{}{confBody: true}, err: config.ErrLoadFailed},
		{params: map[string]interface{}{confBody: "testBody"}, err: nil},

		// Headers
		{params: map[string]interface{}{confHeaders: true}, err: config.ErrLoadFailed},
		{params: map[string]interface{}{confHeaders: "testBody"}, err: config.ErrLoadFailed},
		{params: map[string]interface{}{confHeaders: map[string]string{"test2": "val2"}}, err: config.ErrLoadFailed},
		{params: map[string]interface{}{confHeaders: map[string]interface{}{"test1": "val1"}}, err: nil},

		// Expected code
		{params: map[string]interface{}{confExpectedCode: true}, err: config.ErrLoadFailed},
		{params: map[string]interface{}{confExpectedCode: 200.0}, err: nil},
	}

	check := New()
	for _, tt := range testTable {
		assert.Equal(t, tt.err, check.Init(tt.params))
	}
}

func TestInitLoadedParams(t *testing.T) {
	check := New()
	err := check.Init(map[string]interface{}{
		confMethod:  "get",
		confURL:     "http://test.com/path?k1=v1&k2=v2",
		confBody:    `{"key1":"val1"}`,
		confHeaders: map[string]interface{}{"Authorization": "Bearer test1"}, // The unmarshal will read as map[string]interface{}

		confExpectedCode: 222, // The json unmarshal will read this as float64, but the type switch will work with this too
	})
	require.NoError(t, err)

	req := check.(*httpChecker).req
	assert.Equal(t, http.MethodGet, req.Method)
	assert.Equal(t, "http://test.com/path?k1=v1&k2=v2", req.URL.String())

	b, err := ioutil.ReadAll(req.Body)
	assert.NoError(t, err)
	assert.Equal(t, "Bearer test1", req.Header.Get("authorization"))
	assert.Equal(t, `{"key1":"val1"}`, string(b))
	assert.Equal(t, 222, check.(*httpChecker).expectedCode)
}
