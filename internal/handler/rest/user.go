package rest

import (
	"net/http"
	"twof/blog-api/internal/constant/model"
	"twof/blog-api/internal/constant/state"
	"twof/blog-api/internal/module/user"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	LogUserHandler(*gin.Context)
	RegistrationHandler(*gin.Context)
	Refresh(*gin.Context)
	PasswordReset(c *gin.Context)
	ResetLink(c *gin.Context)
	// GetUsers(c *gin.Context)
}

type userHandler struct {
	userCase user.Usecase
}

func UserInit(userCase user.Usecase) UserHandler {
	return &userHandler{
		userCase,
	}
}
func (u *userHandler) RegistrationHandler(c *gin.Context) {
	ctx := c.Request.Context()
	var user model.User
	if err := c.ShouldBind(&user); err != nil {
		state.ResErr(c, err, http.StatusBadRequest)
		return
	}

	err := u.userCase.Registration(ctx, &user)

	if err != nil {
		state.ResErr(c, err, http.StatusBadRequest)

		return
	}
	user.Password = ""
	state.ResJsonData(c, "User Registered", http.StatusOK, user)
}

func (u *userHandler) LogUserHandler(c *gin.Context) {
	var user model.User

	c.Bind(&user)
	orgUser := user

	err := u.userCase.GetUserByQuery(&user)

	if err != nil {
		state.ResErr(c, err, http.StatusNotFound)

		return
	}

	if isTrue := u.userCase.ComparePassword(user.Password, orgUser.Password); !isTrue {
		state.ResJson(c, "Incorrect Password", http.StatusForbidden)

		return
	}

	token, err := u.userCase.GenerateTokenPair(user)

	if err != nil {
		state.ResJson(c, "Can't generate token", http.StatusForbidden)

	}

	state.ResJsonData(c, "Successfully Loged In", http.StatusOK, token)

}

func (u *userHandler) Refresh(c *gin.Context) {
	tokenReq := model.TokenReqBody{}
	var user model.User

	c.Bind(&tokenReq)

	userId, err := u.userCase.ValidateRefreshToken(tokenReq.RefreshToken)
	if err != nil {
		state.ResErr(c, err, http.StatusBadRequest)

	}

	err = u.userCase.GetUserById(userId, &user)

	if err != nil {
		state.ResErr(c, err, http.StatusBadRequest)
	}

	token, err := u.userCase.GenerateTokenPair(user)

	state.ResJsonData(c, "Successfully Token Refreshed", http.StatusOK, token)

}

func (u *userHandler) ResetLink(c *gin.Context) {

	var resetData model.ResetEmail

	if (c.BindJSON(&resetData)) != nil {
		state.ResJson(c, "Invalid Email", http.StatusBadRequest)
	}

	err := u.userCase.SendResetLink(resetData.Email)

	if err != nil {
		state.ResErr(c, err, http.StatusBadRequest)
	}

	state.ResJson(c, "Email Sent", http.StatusOK)

}

func (u *userHandler) PasswordReset(c *gin.Context) {

	var resetPassword model.PasswordReset

	c.Bind(&resetPassword)
	resetToken, _ := c.GetQuery("reset_token")
	err := u.userCase.ResetPassword(resetPassword, resetToken)

	if err != nil {
		state.ResErr(c, err, http.StatusBadRequest)
	}

	state.ResJson(c, "Password Reseted", http.StatusOK)

}

// func (u userHandler) GetUsers(c *gin.Context) {
// 	var user []model.User

// 	err := u.userCase.GetAllUser(&user)

// 	if err != nil {
// 		c.AbortWithStatus(http.StatusNotFound)
// 	}

// 	c.JSON(http.StatusOK, user)
// }
