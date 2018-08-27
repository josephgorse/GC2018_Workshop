package httpd

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/bmizerany/pat"
	"github.com/gopherguides/training/testing/async/src/keys"
)

// section: handler
// section: handler-struct
type Handler struct {
	Store *keys.Store

	mux *pat.PatternServeMux
}

// section: handler-struct

func NewHandler() *Handler {
	return &Handler{
		mux: pat.New(),
	}
}

// section: handler

// section: servehttp
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mux.Add("POST", "/key", http.HandlerFunc(h.set))
	h.mux.Add("GET", "/key", http.HandlerFunc(h.get))
	h.mux.ServeHTTP(w, r)
}

// section: servehttp

// section: set
func (h *Handler) set(w http.ResponseWriter, r *http.Request) {
	log.Println("set...")
	key := r.FormValue("key")
	value := r.FormValue("value")
	h.Store.Set(key, value)

	w.WriteHeader(http.StatusAccepted)
}

// section: set

// section: get
func (h *Handler) get(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	log.Println("get...")
	key := r.FormValue("key")
	if key == "" {
		http.Error(w, `no key provided`, http.StatusBadRequest)
		return
	}

	value, err := h.Store.Get(key)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response := map[string]interface{}{key: value}

	w.Header().Add("Content-Type", "application/json")
	b, _ := json.Marshal(response)
	w.Write(b)
	log.Printf("took %s", time.Since(now))
}
