package httpd_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gopherguides/training/testing/mocking/src/httpd"
)

func TestSet_NoErrors(t *testing.T) {
	handler := httpd.NewHandler()
	store := &MockStore{}
	store.getFn = func(key string) (interface{}, error) {
		return "bar", nil
	}
	handler.Store = store

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/key", nil)

	r.Form = url.Values{"key": []string{"foo"}, "value": []string{"bar"}}

	handler.ServeHTTP(w, r)
	if exp, got := w.Code, http.StatusAccepted; exp != got {
		t.Errorf("unexpected error code. exp: %d, got %d", exp, got)
	}
}

func TestGet_NoKey(t *testing.T) {
	handler := httpd.NewHandler()
	store := &MockStore{}
	handler.Store = store

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/key", nil)

	handler.ServeHTTP(w, r)
	if exp, got := w.Code, http.StatusBadRequest; exp != got {
		t.Log(w.Body)
		t.Errorf("unexpected error code. exp: %d, got %d", exp, got)
	}
}

func TestGet_NotFound(t *testing.T) {
	handler := httpd.NewHandler()
	store := &MockStore{}
	store.getFn = func(string) (interface{}, error) {
		return nil, notFound{}
	}
	handler.Store = store

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/key?key=foo", nil)

	handler.ServeHTTP(w, r)
	if got, exp := w.Code, http.StatusNotFound; got != exp {
		t.Log(w.Body)
		t.Errorf("unexpected error code. got: %d, exp %d", got, exp)
	}
}

func TestGet_ServerError(t *testing.T) {
	handler := httpd.NewHandler()
	store := &MockStore{}
	store.getFn = func(string) (interface{}, error) {
		return nil, errors.New("boom")
	}
	handler.Store = store

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/key?key=foo", nil)

	handler.ServeHTTP(w, r)
	if got, exp := w.Code, http.StatusInternalServerError; got != exp {
		t.Log(w.Body)
		t.Errorf("unexpected error code. got: %d, exp %d", got, exp)
	}
}

func TestGet_Success(t *testing.T) {
	handler := httpd.NewHandler()
	store := &MockStore{}
	store.getFn = func(string) (interface{}, error) {
		return "bar", nil
	}
	handler.Store = store

	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/key?key=foo", nil)

	handler.ServeHTTP(w, r)
	if got, exp := w.Code, http.StatusOK; got != exp {
		t.Fatalf("unexpected status code.  got %d, expected %d", got, exp)
	}
	data := map[string]interface{}{}
	if err := json.Unmarshal(w.Body.Bytes(), &data); err != nil {
		t.Fatal(err)
	}
	if got, exp := data["foo"], "bar"; got != exp {
		t.Fatalf("unexpected value.  got: %v, exp %v", got, exp)
	}
}

type MockStore struct {
	setFn func(key string, value interface{})
	getFn func(key string) (interface{}, error)
}

func (ms *MockStore) Set(key string, value interface{}) {
	if ms.setFn != nil {
		ms.setFn(key, value)
	}
}

func (ms *MockStore) Get(key string) (interface{}, error) {
	if ms.getFn != nil {
		return ms.getFn(key)
	}
	return nil, nil
}

// not found mock
type notFound struct{}

func (nf notFound) NotFound() {}

func (nf notFound) Error() string {
	return ""
}
