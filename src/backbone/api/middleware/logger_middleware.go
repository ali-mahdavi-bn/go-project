package middleware

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"go-tel/src/backbone/helpers/colors"
	"go-tel/src/backbone/helpers/logger"
	"time"
)

// LoggerMiddleware logs the details of each request.
func LoggerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		currentTime := time.Now()
		err := next(c)
		elapsed := fmt.Sprintf("%.3f", time.Since(currentTime).Seconds())

		req := c.Request()
		res := c.Response()
		var statusCode int
		var messageError string
		if err != nil {
			statusCode = 500
			messageError = err.Error()
		} else {
			statusCode = res.Status
		}
		statusCodeServer := []int{500, 501, 502, 503, 504, 505, 506, 507, 508, 510, 511}

		if err != nil {
			logger.Error(fmt.Sprintf("%s: %s  %s: %s       %s: %s   %s: %s %s: %s",
				colors.GreenColor("Method"),
				req.Method,
				colors.YellowColor("STATUS"),
				colors.RedColor(res.Status),
				colors.MagentaColor("TIME"),
				elapsed,
				colors.PrimaryColor("URL"),
				req.URL.Path,
				colors.RedColor("MESSAGE"),
				colors.RedColor(messageError),
			))
		} else {
			isError := false
			for _, code := range statusCodeServer {
				if statusCode == code {
					isError = true
					break
				}
			}

			if isError {
				logger.Error(fmt.Sprintf("%s: %s  %s: %s       %s: %s   %s: %s",
					colors.GreenColor("Method"),
					req.Method,
					colors.YellowColor("STATUS"),
					colors.RedColor(res.Status),
					colors.MagentaColor("TIME"),
					elapsed,
					colors.PrimaryColor("URL"),
					req.URL.Path))
			} else {
				logger.Info(fmt.Sprintf("%s: %s  %s: %s       %s: %s   %s: %s",
					colors.GreenColor("Method"),
					req.Method,
					colors.YellowColor("STATUS"),
					colors.WhiteColor(res.Status),
					colors.MagentaColor("TIME"),
					elapsed,
					colors.PrimaryColor("URL"),
					req.URL.Path,
				))
			}
		}
		return err
	}
}
