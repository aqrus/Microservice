package middleware

import (
	"net/http"
	"sync"
	"time"
)

type Limiter struct {
	rate  int
	burst int
	mu    sync.Mutex
	// lastRequest is the time of the last request.
	lastRequest time.Time
	// available is the number of requests available to be served.
	available int
}

func NewLimiter(rate, burst int) *Limiter {
	return &Limiter{
		rate:  rate,
		burst: burst,
	}
}

func (l *Limiter) Limit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		l.mu.Lock()
		defer l.mu.Unlock()

		now := time.Now()
		elapsed := now.Sub(l.lastRequest)
		l.lastRequest = now

		// Add tokens to the bucket.
		l.available += int(elapsed.Seconds() * float64(l.rate))
		if l.available > l.burst {
			l.available = l.burst
		}

		// If there are no tokens available, return an error.
		if l.available == 0 {
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
			return
		}

		// Serve the request and decrement the number of available tokens.
		l.available--
		next.ServeHTTP(w, r)
	})
}
