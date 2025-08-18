// Package mux provides support to bind domain level routes
// to the application mux.
package mux

import (
	"net/http"

	"github.com/sralapati/service/apis/services/sales/route/sys/checkapi"
	"github.com/sralapati/service/foundation/logger"
)

// WebAPI constructs a http.Handler with all application routes bound.
func WebAPI(log *logger.Logger) *http.ServeMux {
	mux := http.NewServeMux()

	checkapi.Routes(mux)

	return mux
}
