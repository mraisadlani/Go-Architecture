package middlewares

import (
	"github.com/gin-gonic/gin"
	"go-architecture-mysql/api/exceptions"
	"go-architecture-mysql/api/securities"
)

func SetupAuthenticationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := securities.GetAuthentication(c)

		if err != nil {
			exceptions.AuthorizeException(c, err.Error())
			c.Abort()
			return
		}

		c.Next()
	}
}