package middleware

import (
	"net/http"
	"ranggaAdiPratama/go-with-claude/internal/responses"
	"ranggaAdiPratama/go-with-claude/internal/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "user"
)

func AuthMiddleware(pasetoMaker *utils.PasetoMaker) gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader(authorizationHeaderKey)

		if len(authorizationHeader) == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, responses.Response{
				MetaData: responses.MetaDataResponse{
					Code:    http.StatusUnauthorized,
					Message: "Authorization header is not provided",
				},
			})

			return
		}

		fields := strings.Fields(authorizationHeader)

		if len(fields) < 2 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, responses.Response{
				MetaData: responses.MetaDataResponse{
					Code:    http.StatusUnauthorized,
					Message: "Invalid authorization header format",
				},
			})

			return
		}

		authorizationType := strings.ToLower(fields[0])

		if authorizationType != authorizationTypeBearer {
			c.AbortWithStatusJSON(http.StatusUnauthorized, responses.Response{
				MetaData: responses.MetaDataResponse{
					Code:    http.StatusUnauthorized,
					Message: "Unsupported authorization type",
				},
			})

			return
		}

		accessToken := fields[1]

		payload, err := pasetoMaker.VerifyToken(accessToken)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, responses.Response{
				MetaData: responses.MetaDataResponse{
					Code:    http.StatusUnauthorized,
					Message: "Invalid or expired token",
				},
			})

			return
		}

		c.Set(authorizationPayloadKey, payload)

		c.Next()
	}
}

func RequireRole(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userPayload, exists := c.Get(authorizationPayloadKey)

		if !exists {
			c.AbortWithStatusJSON(http.StatusUnauthorized, responses.Response{
				MetaData: responses.MetaDataResponse{
					Code:    http.StatusUnauthorized,
					Message: "Unauthorized",
				},
			})

			return
		}

		payload, ok := userPayload.(*utils.TokenPayload)

		if !ok {
			c.AbortWithStatusJSON(http.StatusUnauthorized, responses.Response{
				MetaData: responses.MetaDataResponse{
					Code:    http.StatusUnauthorized,
					Message: "Invalid user data",
				},
			})

			return
		}

		hasRole := false

		for _, role := range roles {
			if payload.Role == role {
				hasRole = true
				break
			}
		}

		if !hasRole {
			c.AbortWithStatusJSON(http.StatusForbidden, responses.Response{
				MetaData: responses.MetaDataResponse{
					Code:    http.StatusForbidden,
					Message: "Forbidden: insufficient permissions",
				},
			})

			return
		}

		c.Next()
	}
}
