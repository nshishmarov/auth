package controller

import (
	"encoding/json"
	"fmt"

	"auth/main/dto"
	"auth/main/service/token"

	"github.com/valyala/fasthttp"
)

func Validate(ctx *fasthttp.RequestCtx) {
	body := ctx.Request.Body()
	var templateDTO dto.User
	err := json.Unmarshal(body, &templateDTO)
	if err != nil {
		fmt.Println("Error with parsing data!")
		return
	}

	tokenString, err := token.CreateToken(templateDTO.Username)
	if err != nil {
		fmt.Println("Error with creating token!")
		return
	}
	
	fmt.Println("Token: ", tokenString)
}