package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/majidbigdeli/simpelProject/domin/data"
	"github.com/majidbigdeli/simpelProject/domin/dto"
	"github.com/majidbigdeli/simpelProject/domin/util"
)

//Login ...
func Login(ctx iris.Context) {
	login := dto.Login{}
	if err := ctx.ReadJSON(&login); err != nil {
		resp := dto.NewResponse(false, nil, err.Error())
		ctx.StatusCode(iris.StatusBadRequest)
		var _, _ = ctx.JSON(resp)
		return
	}
	if err := util.IsValid(login); err != nil {
		resp := dto.NewResponse(false, nil, err.Error())
		ctx.StatusCode(iris.StatusBadRequest)
		var _, _ = ctx.JSON(resp)
		return
	}

	user, err := data.GetUserByUserName(login.UserName)

	if err != nil {
		resp := dto.NewResponse(false, nil, err.Error())
		ctx.StatusCode(iris.StatusBadRequest)
		var _, _ = ctx.JSON(resp)
		return
	}

	checkPassword := util.ComparePasswords(user.Password, login.Password)

	if checkPassword == false {
		resp := dto.NewResponse(false, nil, "password is wrong")
		ctx.StatusCode(iris.StatusBadRequest)
		var _, _ = ctx.JSON(resp)
		return
	}

	token, err := util.CreateTokenEndpoint(user)

	if err != nil {
		resp := dto.NewResponse(false, nil, err.Error())
		ctx.StatusCode(iris.StatusInternalServerError)
		var _, _ = ctx.JSON(resp)
		return
	}

	tokenDto := dto.NewToken(token)
	resp := dto.NewResponse(true, tokenDto, "")
	ctx.StatusCode(iris.StatusOK)
	var _, _ = ctx.JSON(resp)
	return
}
