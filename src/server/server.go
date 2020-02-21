package server

import (
	"encoding/json"
	"net/http"

	"github.com/gerifield/mini-monitor/src/cache"
)

type Handler struct {
	listenAddr string
	cache      *cache.Cache
}

func New(listenAddr string, cache *cache.Cache) *Handler {
	return &Handler{
		listenAddr: listenAddr,
		cache:      cache,
	}
}

func (h *Handler) Start() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/api", h.showChecks)

	return http.ListenAndServe(h.listenAddr, mux)
}

func (h *Handler) showChecks(w http.ResponseWriter, r *http.Request) {
	values := h.cache.GetAll()
	json.NewEncoder(w).Encode(struct {
		Checks map[string]bool `json:"checks"`
	}{
		Checks: values,
	})
}
