package apiv1

import (
	"net/http"

	"github.com/sungora/app/core"

	"sample/internal/middle"
)

// PageApiV1 страница api v1
func PageApiV1(w http.ResponseWriter, r *http.Request) {
	var rw = r.Context().Value(middle.KEY_RW).(*core.RW)
	rw.ResponseHtml("PageApiV1", http.StatusOK)
}

