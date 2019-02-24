package middle

import (
	"net/http"
)

// SampleOne middleware
func SampleOne(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(" SampleOne Middle "))
		next.ServeHTTP(w, r)
	})
}

// SampleTwo middleware
func SampleTwo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(" SampleTwo Middle "))
		next.ServeHTTP(w, r)
	})
}

// SampleFour middleware
func SampleFour(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(" SampleFour Middle "))
		next.ServeHTTP(w, r)
	})
}
