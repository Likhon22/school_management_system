package middlewares

import (
	"net/http"
	"school-management-system/pkg/utils"
)

func SetupMiddlewares(mux http.Handler, mw *Middleware) http.Handler {

	return utils.ChainMiddleware(
		mux,
		// mw.Logger,                // log everything including blocked requests
		// mw.Cors,                  // must run early to avoid browser CORS errors
		// mw.IPLimiter.RateLimiter, // block excessive requests before heavy processing
		// mw.SecurityHeaders,       // set headers before sending response
		// mw.Compression,           // compress last, after everything else is wrapped
	)
}
