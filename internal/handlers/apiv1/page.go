package apiv1

import (
	"net/http"
)

// PageApiV1 страница api v1
func PageApiV1(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("PageApiV1"))
}
