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
	token := ctx.Query("token")
	fmt.Println("token from query is:", token)
	req, err := http.NewRequest("POST", "https://timedragon.staging.frimustechnologies.com/v1/auth/check_authenticated", nil)
	if err != nil {
		response.ShowResponse(
			"Error in making request",
			utils.HTTP_BAD_REQUEST,
			"Failure",
			nil,
			ctx,
		)
		ctx.Abort()
		return
	}
	req.Header.Set("Authorization", token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		response.ShowResponse("Server Error", utils.HTTP_INTERNAL_SERVER_ERROR, err.Error(), "", ctx)
		ctx.Abort()
		return
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		response.ShowResponse(
			"Error in reading response body",
			utils.HTTP_BAD_REQUEST,
			"Failure",
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
			"Error in unmarshalling the reponse body",
			utils.HTTP_BAD_REQUEST,
			"Failure",
			nil,
			ctx,
		)
		ctx.Abort()
		return
	}
	fmt.Println("erp details results:", erpDetails)
	ctx.Set("name", erpDetails.Data.Name)
	ctx.Set("emailid", erpDetails.Data.SkypeId.Email)
	ctx.Next()

}
func ResetMiddleware(ctx *gin.Context) {
	var tokenRequest request.ResetPasswordRequest

	err := utils.RequestDecoding(ctx, &tokenRequest)
	if err != nil {
		response.ShowResponse(
			"Error decoding request",
			utils.HTTP_BAD_REQUEST,
			"Failure",
			nil,
			ctx,
		)
		ctx.Abort()
		return
	}

	tokenClaims, err := token.DecodeToken(tokenRequest.Token)
	if err != nil {
		response.ShowResponse(
			"Error decoding login request",
			utils.HTTP_BAD_REQUEST,
			"Failure",
			nil,
			ctx,
		)
		ctx.Abort()
		return
	}

	req, err := http.NewRequest("GET", "https://timedragon.staging.chicmic.co.in/v1/user?_id="+tokenClaims.Id, nil)
	if err != nil {
		response.ShowResponse(
			"Error in making request",
			utils.HTTP_BAD_REQUEST,
			"Failure",
			nil,
			ctx,
		)
		ctx.Abort()
		return
	}
	req.Header.Set("Authorization", tokenRequest.Token)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		response.ShowResponse("Server Error", utils.HTTP_INTERNAL_SERVER_ERROR, err.Error(), "", ctx)
		ctx.Abort()
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		response.ShowResponse(
			"Error in reading response body",
			utils.HTTP_BAD_REQUEST,
			"Failure",
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
			"Error in unmarshalling the reponse body",
			utils.HTTP_BAD_REQUEST,
			"Failure",
			nil,
			ctx,
		)
		ctx.Abort()
		return
	}

	ctx.Set("emailid", erpDetails.Data.SkypeId.Email)
	ctx.Set("password", tokenRequest.Password)
	ctx.Next()

}

func CORSMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", ctx.Request.Header.Get("Origin"))
		// //	ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		// ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		// ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, DELETE, GET, PUT")

		// if ctx.Request.Method != "OPTIONS" {
		// 	ctx.AbortWithStatus(204)
		// 	return
		// }

		//	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", "*")

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(204)
			return
		}

		ctx.Next()
	}
}
