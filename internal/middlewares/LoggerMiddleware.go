package middleware

import (
	"fmt"
	"os"
	"sr-skilltest/internal/model/constant"
	"time"

	"github.com/labstack/echo"
)

// LoggerMiddleware is a custom middleware that logs the requests and responses.
func LoggerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Log the request
		req := c.Request()
		res := c.Response()
		start := time.Now()
		traceID := c.Get(constant.CONTEXT_LOCALS_KEY_TRACE_ID)
		err := next(c)
		if err != nil {
			c.Error(err)
		}
		stop := time.Since(start)

		// Save the logs to file
		f, err := os.OpenFile("prerequest.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			return err
		}
		defer f.Close()

		if _, err = f.WriteString(fmt.Sprintf("%s- %s - %s %s %s %s %d %s\n",
			traceID,
			req.RemoteAddr,
			req.Method,
			req.URL,
			res.Header().Get(echo.HeaderXRequestID),
			req.Body,
			res.Status,
			stop,
		)); err != nil {
			return err
		}

		return nil
	}
}
