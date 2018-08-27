package httpd_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/gopherguides/training/testing/async/src/httpd"
	"github.com/gopherguides/training/testing/async/src/keys"
)

// section: sleep
func TestSet_Sleep(t *testing.T) {
	handler := httpd.NewHandler()
	store := keys.NewStore()
	handler.Store = store

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/key", nil)
	r.Form = url.Values{"key": []string{"foo"}, "value": []string{"bar"}}

	handler.ServeHTTP(w, r)
	if exp, got := w.Code, http.StatusAccepted; exp != got {
		t.Errorf("unexpected error code. exp: %d, got %d", exp, got)
	}

	w = httptest.NewRecorder()
	r = httptest.NewRequest("GET", "/key?key=foo", nil)

	time.Sleep(4 * time.Second)

	handler.ServeHTTP(w, r)
	if exp, got := w.Code, http.StatusOK; exp != got {
		t.Log(w.Body)
		t.Errorf("unexpected error code. exp: %d, got %d", exp, got)
	}
}

// section: sleep

// section: channels
func TestSet_Channels(t *testing.T) {
	handler := httpd.NewHandler()
	store := keys.NewStore()
	handler.Store = store

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/key", nil)
	r.Form = url.Values{"key": []string{"foo"}, "value": []string{"bar"}}

	handler.ServeHTTP(w, r)
	if exp, got := w.Code, http.StatusAccepted; exp != got {
		t.Errorf("unexpected error code. exp: %d, got %d", exp, got)
	}

	test := func() error {
		w = httptest.NewRecorder()
		r = httptest.NewRequest("GET", "/key?key=foo", nil)
		handler.ServeHTTP(w, r)

		if got, exp := w.Code, http.StatusOK; got != exp {
			return fmt.Errorf("unexpected status code.  got %d, expected %d", got, exp)
		}
		data := map[string]interface{}{}
		if err := json.Unmarshal(w.Body.Bytes(), &data); err != nil {
			return err
		}
		if got, exp := data["foo"], "bar"; got != exp {
			return fmt.Errorf("unexpected value.  got: %v, exp %v", got, exp)
		}
		// test successful
		return nil
	}

	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()
	timeout := time.NewTimer(5 * time.Second)
	defer timeout.Stop()

	var testErr error

	for {
		select {
		case <-timeout.C:
			t.Fatalf("test timed out waiting for success.  last error: %s", testErr)
			return
		case <-ticker.C:
			testErr = test()
			if testErr == nil {
				// test successful
				return
			}
		}
	}
}

// section: channels

// section: foo
func TestFoo(t *testing.T) {
	handler := httpd.NewHandler()
	store := keys.NewStore()
	handler.Store = store

	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/key", nil)
	r.Form = url.Values{"key": []string{"foo"}, "value": []string{"bar"}}

	handler.ServeHTTP(w, r)
	if exp, got := w.Code, http.StatusAccepted; exp != got {
		t.Errorf("unexpected error code. exp: %d, got %d", exp, got)
	}

	test := func() error {
		w = httptest.NewRecorder()
		r = httptest.NewRequest("GET", "/key?key=foo", nil)
		handler.ServeHTTP(w, r)

		if got, exp := w.Code, http.StatusOK; got != exp {
			return fmt.Errorf("unexpected status code.  got %d, expected %d", got, exp)
		}
		data := map[string]interface{}{}
		if err := json.Unmarshal(w.Body.Bytes(), &data); err != nil {
			return err
		}
		if got, exp := data["foo"], "bar"; got != exp {
			return fmt.Errorf("unexpected value.  got: %v, exp %v", got, exp)
		}
		// test successful
		return nil
	}

	// section: timeout-example
	if err := TimeoutAfter(5*time.Second, func() error {
		if err := test(); err != nil {
			return err
		}
		return nil
	}); err != nil {
		t.Fatal(err)
	}
	// section: timeout-example
}

// section: foo

// section: timeout-after
// TimeoutAfter returns an error if fn doesn't return a nil error within the timeout duration.
func TimeoutAfter(timeout time.Duration, fn func() error) error {
	timer := time.NewTimer(timeout)
	defer timer.Stop()

	ticker := time.NewTicker(10 * time.Millisecond)
	defer ticker.Stop()

	var err error
	for {
		// Run the function and save the last error, if any. Exit if no error.
		if err = fn(); err == nil {
			return nil
		}

		// Fail test if timeout occurs.
		// Otherwise wait for initial channel or interval channel.
		select {
		case <-timer.C:
			return fmt.Errorf("%s (%s timeout)", err, timeout)
		case <-ticker.C:
			continue
		}
	}
}
