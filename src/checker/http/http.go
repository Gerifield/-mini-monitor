package http

import (
	"net/http"
	"strings"

	"github.com/Gerifield/mini-monitor/src/config"
	"github.com/Gerifield/mini-monitor/src/loader"
)

const (
	confMethod       = "method"
	confURL          = "url"
	confBody         = "body"
	confHeaders      = "headers"
	confExpectedCode = "expectedCode"
)

type httpChecker struct {
	client       *http.Client
	req          *http.Request
	expectedCode int
}

func New() config.Checker {
	return &httpChecker{}
}

func (h *httpChecker) Init(conf map[string]interface{}) error {

	// Load the method
	method, err := loader.ConfigString(conf, confMethod)
	if err != nil {
		return err
	}
	method = strings.ToUpper(method)

	// Load the URL
	url, err := loader.ConfigString(conf, confURL)
	if err != nil {
		return err
	}

	// Load the body
	body, err := loader.ConfigString(conf, confBody)
	if err != nil {
		return err
	}

	// Load the headers
	headers := make(map[string]string)
	if h, ok := conf[confHeaders]; ok {
		if headersMap, ok := h.(map[string]interface{}); ok {
			for k, v := range headersMap {
				if valStr, ok := v.(string); !ok {
					return config.ErrLoadFailed
				} else {
					headers[k] = valStr
				}
			}
		} else {
			return config.ErrLoadFailed
		}
	}

	// Load the expected code
	h.expectedCode, err = loader.ConfigInt(conf, confExpectedCode)
	if err != nil {
		return err
	}

	h.client = http.DefaultClient // TODO: Make this configurable too

	// Set the HTTP request
	h.req, err = http.NewRequest(method, url, strings.NewReader(body))
	if err != nil {
		return err
	}

	for k, v := range headers {
		h.req.Header.Set(k, v)
	}
	return nil
}

func (h *httpChecker) Check() error {
	resp, err := h.client.Do(h.req)
	if err != nil {
		return err
	}

	if resp.StatusCode != h.expectedCode {
		return config.ErrCheckFailed
	}
	return nil
}
