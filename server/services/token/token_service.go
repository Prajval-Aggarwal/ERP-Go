package token

import (
	"fmt"
	"main/server/model"
	"main/server/response"
	"main/server/utils"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// Generate JWT Token
func GenerateToken(claims model.Claims, ctx *gin.Context) *string {
	//create user claims

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(os.Getenv("JWTKEY")))

	if err != nil {
		response.ShowResponse("Error signing token", utils.HTTP_UNAUTHORIZED, "Failure", nil, ctx)
		return nil
	}
	return &tokenString
}

// Decode Token function
func DecodeToken(tokenString string) (*model.Claims, error) {
	claims := &model.Claims{}
	parsedToken, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("error")
		}
		return []byte(os.Getenv("JWTKEY")), nil
	})

	if err != nil || !parsedToken.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}
