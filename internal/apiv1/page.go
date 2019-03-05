package apiv1

import (
	"net/http"

	"github.com/sungora/app/core"
	"github.com/sungora/app/servhttp/middleware"
)

// PageApiV1 страница api v1
func PageApiV1(w http.ResponseWriter, r *http.Request) {
	var rw = r.Context().Value(middleware.KeyRW).(*core.RW)
	rw.ResponseHtml("PageApiV1", http.StatusOK)
}
