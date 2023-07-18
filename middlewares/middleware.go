package middlewares

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
)

// logging middleware
func LoggingMiddleware(nxt http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Before handler is executed")
		w.Write([]byte("Adding response via middleware\n"))
		log.Println(r.URL.Path)
		nxt.ServeHTTP(w, r)
		fmt.Println("After handler is executed")
	})
}

// adding headers middleware
func HeaderMiddleware(nxt http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		nxt(w, r)
	}
}

// recovery middleware
func PanicRecoveryMiddleware(nxt http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				log.Println(string(debug.Stack()))
			}
		}()
		nxt(w, r)
	}
}

// cors middleware
func CORSMiddleware(nxt http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		w.Header().Set("Access-Control-Allow-Origin", origin)
		if r.Method == "OPTIONS" {
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "GET,POST,DELETE,PUT")
			w.Header().Set("Access-Control-Allow-Headers", "Control-Type, X-CSRF-Token, Authorization")
			return
		} else {
			nxt.ServeHTTP(w, r)
		}
	})
}

//chaining the middlewares
type Middleware func(http.Handler) http.Handler

func ChainingMiddleware(h http.Handler, m ...Middleware) http.Handler {
	if len(m) < 1 {
		return h
	}

	wrappedHandler := h
	for i := len(m) - 1; i >= 0; i-- {
		wrappedHandler = m[i](wrappedHandler)
	}

	return wrappedHandler
}
