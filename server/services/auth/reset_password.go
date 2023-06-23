package auth

import (
	"fmt"
	"main/server/db"
	"main/server/response"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func ResetMattermostPassword(ctx *gin.Context, password string, emailId string) {
	fmt.Println("passwrod", password)
	fmt.Println("email", emailId)

	//hashedPass := HashPassword(password)
	// query := "UPDATE users SET password = '" + password + "' WHERE email = '" + emailId + "'"
	// fmt.Println("query", query)
	query := "UPDATE users SET password=? WHERE email=?"
	db.RawExecutor(query, password, emailId)
	response.ShowResponse("Password update sucessfully", 200, "Success", nil, ctx)

}

func HashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		panic(err)
	}

	return string(hash)
}
