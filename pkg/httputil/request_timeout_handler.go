package httputil

import (
	"context"
	"net/http"
	"sync"
	"time"
)

type RequestTimeoutHandler struct {
	handler        http.Handler
	defaultTimeout time.Duration

	pathTimeoutMapLock sync.RWMutex
	pathTimeoutMap     map[string]time.Duration
}

func (rth *RequestTimeoutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	timeout := rth.defaultTimeout

	if r.URL != nil {
		rth.pathTimeoutMapLock.RLock()
		timeoutByPath, isExist := rth.pathTimeoutMap[r.URL.Path]
		rth.pathTimeoutMapLock.RUnlock()

		if isExist {
			timeout = timeoutByPath
		}
	}

	contextWithTimeout, cancel := context.WithTimeout(r.Context(), timeout)
	defer cancel()

	requestWithContext := r.WithContext(contextWithTimeout)

	rth.handler.ServeHTTP(w, requestWithContext)
}

func (rth *RequestTimeoutHandler) SetPathTimeoutMap(pathTimeoutMap map[string]time.Duration) {
	rth.pathTimeoutMapLock.Lock()
	rth.pathTimeoutMap = pathTimeoutMap
	rth.pathTimeoutMapLock.Unlock()
}

func NewRequestTimeout(handler http.Handler, defaultTimeout time.Duration, pathTimeoutMap map[string]time.Duration) http.Handler {
	return &RequestTimeoutHandler{
		handler:        handler,
		defaultTimeout: defaultTimeout,
		pathTimeoutMap: pathTimeoutMap,
	}
}
