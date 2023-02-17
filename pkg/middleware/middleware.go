package middleware

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/ztrue/tracerr"
	"net/http"
)

// PanicHandler provides middleware which could recover panic, print stack trace and inc counter
func PanicHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				err, ok := rec.(error)
				if !ok {
					err = fmt.Errorf("%v", rec)
				}

				e := tracerr.Wrap(err)
				frames := e.StackTrace()[2:]

				tracerr.PrintSourceColor(tracerr.CustomError(err, frames))

				w.WriteHeader(http.StatusInternalServerError)

				PanicCounter.Inc()
			}
		}()

		handler.ServeHTTP(w, r)
	})
}

// RequestsCountHandler middleware for accounting HTTP requests
func RequestsCountHandler(handler http.Handler) http.Handler {
	return promhttp.InstrumentHandlerCounter(RequestsCounter, handler)
}

// UnauthorizedRequestsCountHandler middleware for accounting HTTP requests with empty X-Token
func UnauthorizedRequestsCountHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.Header.Get("X-Token") == "" {
			UnauthorizedRequestsCounter.Inc()
		}
		handler.ServeHTTP(writer, request)
	})
}

// RequestsDurationHandler middleware for grab requests duration
func RequestsDurationHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		timer := prometheus.NewTimer(RequestDuration)
		handler.ServeHTTP(writer, request)
		timer.ObserveDuration()
	})
}
