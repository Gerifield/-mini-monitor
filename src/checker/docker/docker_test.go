package docker

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gerifield/mini-monitor/src/config"
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

func TestDockerInitDebug(t *testing.T) {
	testTable := []struct {
		params map[string]interface{}
		err    error
	}{
		{params: nil, err: nil},
		{params: map[string]interface{}{}, err: nil},
		{params: map[string]interface{}{"testKey1": "testVal1"}, err: nil},
		{params: map[string]interface{}{confDebug: "testRegex"}, err: config.ErrLoadFailed},
		{params: map[string]interface{}{confDebug: true}, err: nil},
	}

	check := New()
	for _, tt := range testTable {
		assert.Equal(t, tt.err, check.Init(tt.params))
	}
}

func TestDoChecks(t *testing.T) {
	testTable := []struct {
		input []dockerPSOutput
		err   error
	}{
		{nil, config.ErrCheckFailed},
		{[]dockerPSOutput{genOutput("noID", "noImage", "noName")}, config.ErrCheckFailed},
		{[]dockerPSOutput{genOutput("testID", "noImage", "noName")}, nil},
		{[]dockerPSOutput{genOutput("noID", "testImage", "noName")}, nil},
		{[]dockerPSOutput{genOutput("noID", "noImage", "testName")}, nil},
	}

	check := New()
	assert.NoError(t, check.Init(map[string]interface{}{
		confID:         "testID",
		confImageRegex: "testImage",
		confNameRegex:  "testName",
	}))

	for _, tt := range testTable {
		assert.Equal(t, tt.err, check.(*dockerChecker).doChecks(tt.input))
	}
}

func genOutput(ID, image, name string) dockerPSOutput {
	return dockerPSOutput{
		ID:    ID,
		Image: image,
		Names: name,
	}
}
