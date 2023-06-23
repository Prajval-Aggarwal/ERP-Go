package auth

import (
	"fmt"
	"main/server/response"
	"main/server/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LogoutService(ctx *gin.Context) {

	//deleting the cookies
	mmauthCookie := &http.Cookie{
		Name:     "MMAUTHTOKEN",
		Value:    "",
		MaxAge:   -1,
		Path:     "/",
		HttpOnly: false,
	}
	mmuserCookie := &http.Cookie{
		Name:     "MMUSERID",
		Value:    "",
		MaxAge:   -1,
		Path:     "/",
		HttpOnly: false,
	}

	mmcsrfCookie := &http.Cookie{
		Name:     "MMCSRF",
		Value:    "",
		MaxAge:   -1,
		Path:     "/",
		HttpOnly: false,
	}

	http.SetCookie(ctx.Writer, mmauthCookie)
	http.SetCookie(ctx.Writer, mmuserCookie)
	http.SetCookie(ctx.Writer, mmcsrfCookie)

	//call mattermost logout api
	reqst, err := http.NewRequest("POST", "http://192.180.0.123:8065/api/v4/users/logout", nil)

	if err != nil {
		response.ShowResponse("Server Error", utils.HTTP_INTERNAL_SERVER_ERROR, err.Error(), nil, ctx)
		return

	}
	reqst.Header.Add("X-Requested-With", "XMLHttpRequest")
	//reqst.Header.Add("Origin", ctx.Request.Header.Get("Origin"))
	fmt.Println("request: ", reqst)
	resp, err := http.DefaultClient.Do(reqst)
	if err != nil {
		response.ShowResponse("Server Error", utils.HTTP_INTERNAL_SERVER_ERROR, err.Error(), nil, ctx)
		return
	}
	fmt.Println("response is:", resp)
	if resp.StatusCode != 200 {
		response.ShowResponse("Error", int64(resp.StatusCode), "", nil, ctx)
		return
	}
	response.ShowResponse("Logout sucessfully", 200, "Success", nil, ctx)
}
