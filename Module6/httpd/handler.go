package httpd

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/bmizerany/pat"
)

type Handler struct {
	Store interface {
		Set(key string, value interface{})
		Get(key string) (interface{}, error)
	}

	mux *pat.PatternServeMux
}

func NewHandler() *Handler {
	return &Handler{
		mux: pat.New(),
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mux.Add("POST", "/key", http.HandlerFunc(h.set))
	h.mux.Add("GET", "/key", http.HandlerFunc(h.get))
	h.mux.ServeHTTP(w, r)
}

func (h *Handler) set(w http.ResponseWriter, r *http.Request) {
	log.Println("set...")
	key := r.FormValue("key")
	value := r.FormValue("value")
	h.Store.Set(key, value)

	w.WriteHeader(http.StatusAccepted)
}

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
		if isNotFound(err) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{key: value}

	w.Header().Add("Content-Type", "application/json")
	b, _ := json.Marshal(response)
	w.Write(b)
	log.Printf("took %s", time.Since(now))
}

func isNotFound(err error) bool {
	type notFound interface {
		NotFound()
	}
	_, ok := err.(notFound)
	return ok
}
