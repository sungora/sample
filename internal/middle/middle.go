package middle

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/sungora/app/core"
)

// SampleOne middleware
func SampleOne(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(" SampleOne Middle ")
		next.ServeHTTP(w, r)
	})
}

// SampleTwo middleware
func SampleTwo(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(" SampleTwo Middle ")
		next.ServeHTTP(w, r)
	})
}

// SampleFour middleware
func SampleFour(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(" SampleFour Middle ")
		next.ServeHTTP(w, r)
	})
}

const KEY_RW string = "RW"

// Main (middleware)
// инициализация таймаута контекста для контроля времени выполениня запроса
// инициализация инструмента для ответа
func Main(d time.Duration) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx, cancel := context.WithTimeout(r.Context(), d)
			defer cancel()
			next.ServeHTTP(w, r.WithContext(context.WithValue(ctx, KEY_RW, core.NewRW(w, r))))
		})
	}
}

// NotFound
func NotFound(w http.ResponseWriter, r *http.Request) {
	var rw = r.Context().Value(KEY_RW).(*core.RW)
	rw.ResponseStatic(core.Config.DirWww + r.URL.Path)
}
