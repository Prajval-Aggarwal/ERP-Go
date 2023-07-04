package auth

import (
	"fmt"
	"main/server/db"
	"main/server/response"
	"main/server/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func ResetMattermostPassword(ctx *gin.Context, password string, emailId string) {
	fmt.Println("password", password)
	fmt.Println("email", emailId)

	//hashedPass := HashPassword(password)
	// query := "UPDATE users SET password = '" + password + "' WHERE email = '" + emailId + "'"
	// fmt.Println("query", query)
	query := "UPDATE users SET password=? WHERE email=?"
	err := db.RawExecutor(query, password, emailId)
	if err != nil {
		response.ShowResponse(utils.QUERYEXECUTOR_ERROR, utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		return
	}
	response.ShowResponse(utils.PASSWORD_UPDATE_SUCCESSFULL, utils.HTTP_OK, utils.SUCCESS, nil, ctx)

}

func HashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		panic(err)
	}

	return string(hash)
}
