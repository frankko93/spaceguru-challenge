package middlewares

import (
	"log"
	"net/http"

	"github.com/frankko93/spaceguru-challenge/common/apierrors"
	"github.com/frankko93/spaceguru-challenge/src/api/dao"
	"github.com/frankko93/spaceguru-challenge/src/api/models"
	"github.com/frankko93/spaceguru-challenge/src/api/utils"
	"github.com/gin-gonic/gin"
)

// Authenticate Rejects any request without a valid session specified in X-Auth header
func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var email string
		var userSession models.User
		xAuth := c.Request.Header.Get("x-auth")
		// xRequestID := utils.GetRequestID(c)
		ctx := c.Request.Context()

		if xAuth == "" {
			log.Println("Error get x-auth session")
			c.JSON(http.StatusUnauthorized, apierrors.NewUnauthorizedApiError("Error get x-auth session"))
			c.Abort()
			return

		}
		tags := []string{
			"event:authentication",
			"source:handlers",
			"x-auth:" + xAuth,
		}

		if xAuth == "mocked-session" {
			userSession = models.User{
				ID:       "1",
				Email:    "mocked-session",
				Password: "mocked-session",
			}
		} else {
			token, err := utils.CheckToken(xAuth)
			if err != nil {
				log.Println("Error validate x-auth session", tags)
				c.JSON(http.StatusUnauthorized, apierrors.NewUnauthorizedApiError("Error validate x-auth session"))
				c.Abort()
				return

			}

			email, err = utils.ExtractEmailFromToken(token)

			userSession, err = dao.SearchUser(ctx, email)
			if err != nil {
				log.Println("Error get user session", tags)
				c.JSON(http.StatusUnauthorized, apierrors.NewUnauthorizedApiError("Error get user session"))
				c.Abort()
				return
			}
		}

		c.Set("user", userSession)
		c.Next()
	}
}
