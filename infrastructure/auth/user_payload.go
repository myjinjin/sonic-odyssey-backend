package auth

import (
	"encoding/json"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

func GetUserPayload(c *gin.Context, md *jwt.GinJWTMiddleware) *UserPayload {
	claims, jwtErr := md.GetClaimsFromJWT(c)
	if jwtErr == nil {
		result := claims[md.IdentityKey]
		if result == nil {
			return nil
		}
		data := result.(map[string]interface{})

		jsonString, _ := json.Marshal(data)

		payload := UserPayload{}
		err := json.Unmarshal(jsonString, &payload)
		if err != nil {
			return nil
		}
		return &payload
	}
	return nil
}
