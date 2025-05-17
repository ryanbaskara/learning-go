package readiness

import (
	"fmt"
	"net/http"
	"sync"
)

// Ready function for readiness probe
type Ready struct {
	cb     chan bool
	status bool
	mutex  sync.Mutex
}

// NewReady build ready
func NewReady() *Ready {
	r := &Ready{
		cb:     make(chan bool),
		status: false,
	}

	go r.monitor()

	return r
}

// Handler is used to control the flow of GET /ready endpoint
func (rd *Ready) Handler(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" && r.URL.Path == "/ready" {
			rd.ServeHTTP(w, r)

			return
		}

		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

// ServeHTTP serve http request for readiness state
func (rd *Ready) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rd.mutex.Lock()
	defer rd.mutex.Unlock()
	if rd.status {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "ok")

		return
	}

	w.WriteHeader(http.StatusBadGateway)
	fmt.Fprintln(w, "not ok")
}

func (rd *Ready) monitor() {
	for b := range rd.cb {
		rd.mutex.Lock()
		rd.status = b
		rd.mutex.Unlock()
	}
}

func (rd *Ready) Resume() {
	rd.cb <- true
}

func (rd *Ready) Stop() {
	rd.cb <- false
}
