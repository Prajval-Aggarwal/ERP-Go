package gateway

import (
	"encoding/json"
	"fmt"
	"io"
	"main/server/model"
	"main/server/request"
	"main/server/response"
	"main/server/services/token"
	"main/server/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserDetailsMiddleware(ctx *gin.Context) {
	var token string
	if ctx.Query(utils.TOKEN) != "" {
		token = ctx.Query(utils.TOKEN)
	} else if ctx.Request.Header.Get(utils.AUTHORIZATION_HEADER) != "" {
		token = ctx.Request.Header.Get(utils.AUTHORIZATION_HEADER)
	}
	fmt.Println("token from query is:", token)
	if token == "" {
		response.ShowResponse(utils.TOKEN_NOT_FOUND, utils.HTTP_BAD_REQUEST, utils.FAILURE, nil, ctx)
		ctx.Abort()
		return
	}
	req, err := http.NewRequest(utils.REQUEST_POST, utils.STAGING_USER_AUTHENTICATION_URL, nil)
	if err != nil {
		response.ShowResponse(
			utils.ERROR_IN_HTTP_REQUEST,
			utils.HTTP_BAD_REQUEST,
			utils.FAILURE,
			nil,
			ctx,
		)
		ctx.Abort()
		return
	}
	req.Header.Set(utils.AUTHORIZATION_HEADER, token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		response.ShowResponse(utils.SERVER_ERROR, utils.HTTP_INTERNAL_SERVER_ERROR, err.Error(), "", ctx)
		ctx.Abort()
		return
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		response.ShowResponse(
			utils.RESPONSE_BODY_ERROR,
			utils.HTTP_BAD_REQUEST,
			utils.FAILURE,
			nil,
			ctx,
		)
		ctx.Abort()
		return
	}

	var erpDetails model.ErpDetails
	err = json.Unmarshal(body, &erpDetails)
	if err != nil {
		response.ShowResponse(
			utils.UNMARSHALLING_ERROR,
			utils.HTTP_BAD_REQUEST,
			utils.FAILURE,
			nil,
			ctx,
		)
		ctx.Abort()
		return
	}
	fmt.Println("erp details results:", erpDetails)
	ctx.Set("name", erpDetails.Data.EmployeeId)
	ctx.Set("emailid", erpDetails.Data.SkypeId.Email)
	ctx.Next()

}
func ResetMiddleware(ctx *gin.Context) {
	var tokenRequest request.ResetPasswordRequest

	err := utils.RequestDecoding(ctx, &tokenRequest)
	if err != nil {
		response.ShowResponse(
			utils.DECODING_ERROR,
			utils.HTTP_BAD_REQUEST,
			utils.FAILURE,
			nil,
			ctx,
		)
		ctx.Abort()
		return
	}

	tokenClaims, err := token.DecodeToken(tokenRequest.Token)
	if err != nil {
		response.ShowResponse(
			utils.DECODING_ERROR,
			utils.HTTP_BAD_REQUEST,
			utils.FAILURE,
			nil,
			ctx,
		)
		ctx.Abort()
		return
	}

	req, err := http.NewRequest(utils.REQUEST_GET, utils.STAGING_USER_URL+tokenClaims.Id, nil)
	if err != nil {
		response.ShowResponse(
			utils.ERROR_IN_HTTP_REQUEST,
			utils.HTTP_BAD_REQUEST,
			utils.FAILURE,
			nil,
			ctx,
		)
		ctx.Abort()
		return
	}
	req.Header.Set(utils.AUTHORIZATION_HEADER, tokenRequest.Token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		response.ShowResponse(utils.SERVER_ERROR, utils.HTTP_INTERNAL_SERVER_ERROR, err.Error(), "", ctx)
		ctx.Abort()
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		response.ShowResponse(
			utils.RESPONSE_BODY_ERROR,
			utils.HTTP_BAD_REQUEST,
			utils.FAILURE,
			nil,
			ctx,
		)
		ctx.Abort()
		return
	}

	var erpDetails model.ErpDetails
	err = json.Unmarshal(body, &erpDetails)
	if err != nil {
		response.ShowResponse(
			utils.UNMARSHALLING_ERROR,
			utils.HTTP_BAD_REQUEST,
			utils.FAILURE,
			nil,
			ctx,
		)
		ctx.Abort()
		return
	}
	fmt.Println("ERP details: ", erpDetails)
	ctx.Set("emailid", erpDetails.Data.SkypeId.Email)
	ctx.Set("password", tokenRequest.Password)
	ctx.Next()

}

func CORSMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("origin is", ctx.Request.Header.Get("Origin"))
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", ctx.Request.Header.Get("Origin"))
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, DELETE, GET, PUT")

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(204)
			return
		}

		ctx.Next()
	}
}
