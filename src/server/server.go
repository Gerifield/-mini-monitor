package server

import (
	"encoding/json"
	"html/template"
	"net/http"

	"github.com/Gerifield/mini-monitor/src/cache"
)

type Handler struct {
	listenAddr string
	cache      *cache.Cache

	mainTemplate *template.Template
}

func New(listenAddr string, cache *cache.Cache) *Handler {
	return &Handler{
		listenAddr: listenAddr,
		cache:      cache,

		mainTemplate: template.Must(template.ParseFiles("src/server/template/index.tpl")),
	}
}

func (h *Handler) Start() error {
	mux := http.NewServeMux()

	mux.HandleFunc("/", h.renderTemplate)
	mux.HandleFunc("/api", h.showChecks)

	return http.ListenAndServe(h.listenAddr, mux)
}

func (h *Handler) renderTemplate(w http.ResponseWriter, r *http.Request) {
	_ = h.mainTemplate.Execute(w, struct {
		Values map[string]bool
	}{
		Values: h.cache.GetAll(),
	})
}

func (h *Handler) showChecks(w http.ResponseWriter, r *http.Request) {
	values := h.cache.GetAll()
	_ = json.NewEncoder(w).Encode(struct {
		Checks map[string]bool `json:"checks"`
	}{
		Checks: values,
	})
}
