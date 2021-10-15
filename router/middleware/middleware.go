package middleware

import (
	"time"

	"github.com/labstack/echo-contrib/prometheus"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const (
	// Key (Should come from somewhere else).
	Key = "secret"
)

func JWTSkipper(c echo.Context) bool {
	// skippedUris := []string{"/login", "/signup", "/healthcheck", "/metrics"} // URIs to skip JWT validation
	// for _, uri := range skippedUris {
	// 	if c.Path() == uri {
	// 		return true
	// 	}
	// }
	return false
}

func SetMainMiddlewares(e *echo.Echo) {
	// middlewares
	// logger
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		// Common Log Format - 127.0.0.1 PostmanRuntime/7.28.0 - [13/Oct/2021 12:50:13 +0800] "GET /api/customers/1 HTTP/1.1" 200 190
		Format:           "${remote_ip} ${user_agent} - [${time_custom}] \"${method} ${path} ${protocol}\" ${status} ${bytes_out}" + "\n",
		CustomTimeFormat: "02/Jan/2006 15:04:05 -0700",
	}))

	// timeout
	e.Use(middleware.TimeoutWithConfig(middleware.TimeoutConfig{
		Timeout: 5 * time.Second, // 5 seconds timeout
	}))

	// cors
	e.Use(middleware.CORS())

	// recover from panic
	e.Use(middleware.Recover())

	// prometheus metrics - visit <host>:<port>/metrics
	p := prometheus.NewPrometheus("echo", nil)
	p.Use(e)
}

func SetApiMiddlewares(g *echo.Group) {
	// jwt
	g.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		Skipper:       JWTSkipper,
		SigningMethod: "HS512", // does not work when changed to HS256
		SigningKey:    []byte(Key),
		TokenLookup:   "cookie:JWTCookie",
	}))
}
