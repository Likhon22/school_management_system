package middlewares

import "net/http"

func ChainMiddleware(mux http.Handler, mws ...func(http.Handler) http.Handler) http.Handler {

	for _, mw := range mws {
		mux = mw(mux)

	}
	return mux
}
