package controllers

import "github.com/gin-gonic/gin"

// ErrorMiddleware returns a middleware that checks if there's been an error in
// the request and returns its message.
func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		errors := c.Errors.Errors()
		if len(errors) > 0 {
			c.String(c.Writer.Status(), errors[0])
			c.Abort()
		}
	}
}
