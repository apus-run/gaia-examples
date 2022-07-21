package pkg

import (
	"context"
	"github.com/apus-run/gaia/middleware"
	"github.com/gin-gonic/gin"
)

func Middlewares(m ...middleware.Middleware) gin.HandlerFunc {
	chain := middleware.Chain(m...)
	return func(c *gin.Context) {
		next := func(ctx context.Context, req interface{}) (interface{}, error) {
			c.Next()
			return c.Writer, nil
		}
		next = chain(next)
		next(c.Request.Context(), c.Request)
	}
}
