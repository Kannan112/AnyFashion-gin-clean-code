package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	handlerUtil "github.com/kannan112/go-gin-clean-arch/pkg/api/handlerUril"
	"github.com/kannan112/go-gin-clean-arch/pkg/common/req"
	"github.com/kannan112/go-gin-clean-arch/pkg/common/res"
	services "github.com/kannan112/go-gin-clean-arch/pkg/usecase/interface"
)

type UserHandler struct {
	userUseCase services.UserUseCase
	cartUseCase services.CartUseCases
}

func NewUserHandler(usecase services.UserUseCase) *UserHandler {
	return &UserHandler{
		userUseCase: usecase,
	}
}

//---------------------------------_UserSignUp-----------------------------

func (cr *UserHandler) UserSignUp(c *gin.Context) {
	var user req.UserReq
	err := c.Bind(&user)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, res.Response{
			StatusCode: 422,
			Message:    "can't bind",
			Data:       nil,
			Errors:     err.Error(),
		})
	}
	userData, err := cr.userUseCase.UserSignUp(c.Request.Context(), user)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "unable signup",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	err = cr.cartUseCase.CreateCart(userData.Id)

	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "unable create cart",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, res.Response{
		StatusCode: 201,
		Message:    "user signup Scart_itemsuccessfully",
		Data:       userData,
		Errors:     nil,
	})

}

// -------------------------------UserLogin-------------------

// LoginWithEmail
// @Summary User Login
// @ID UserLogin
// @Description Login as a user to access the ecommerce site
// @Tags Users
// @Accept json
// @Produce json
// @Param   input   body     req.LoginReq{}   true  "Input Field"
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Router /user/login [post]
func (cr *UserHandler) UserLogin(c *gin.Context) {
	var user req.LoginReq
	err := c.Bind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "failed to read request body",
			Data:       nil,
			Errors:     err.Error(),
		})
		return

	}
	sessionValue, err := cr.userUseCase.UserLogin(c.Request.Context(), user)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "failed to login",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("UserAuth", sessionValue, 3600*24*30, "", "", false, true)
	c.JSON(http.StatusOK, res.Response{
		StatusCode: 200,
		Message:    "logined successfuly",
		Data:       nil,
		Errors:     nil,
	})

}

// -------------------------------UserLogout-------------------

// Logout
// @Summary User Logout
// @ID UserLogout
// @Description User logout to access the ecommerce site
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {object} res.Response
// @Failure 400 {object} res.Response
// @Router /user/logout [get]
func (cr *UserHandler) UserLogout(c *gin.Context) {
	c.SetCookie("UserAuth", "", 1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{
		"message": "UserLogouted",
	})
}
func (cr *UserHandler) AddAddress(c *gin.Context) {
	id, err := handlerUtil.GetUserIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "failed to get the userId",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	var address req.Address
	err = c.Bind(&address)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "failed to bind Address",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	err = cr.userUseCase.AddAddress(id, address)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "Can't add address",
			Data:       nil,
			Errors:     err.Error(),
		})
		return

	}
	c.JSON(http.StatusAccepted, res.Response{
		StatusCode: 200,
		Message:    "address added successfully",
		Data:       address,
		Errors:     nil,
	})

}
func (cr *UserHandler) UpdateAddress(c *gin.Context) {
	paramsId := c.Param("addressId")
	addressId, err := strconv.Atoi(paramsId)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "Can't find ProductId",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	id, err := handlerUtil.GetUserIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "failed to get the userId",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	var address req.Address
	err = c.Bind(&address)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "failed to bind Address",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	err = cr.userUseCase.UpdateAddress(id, addressId, address)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "Can't add address",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusAccepted, res.Response{
		StatusCode: 200,
		Message:    "address updated successfully",
		Data:       address,
		Errors:     nil,
	})
}

// list all addresses of user
func (cr *UserHandler) ListallAddress(c *gin.Context) {
	id, err := handlerUtil.GetUserIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "failed to get the userId",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	data, err := cr.userUseCase.ListallAddress(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "failed to list",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusAccepted, res.Response{
		StatusCode: 200,
		Message:    "All Address",
		Data:       data,
		Errors:     nil,
	})
}

func (cr *UserHandler) ViewProfile(c *gin.Context) {
	id, err := handlerUtil.GetUserIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "failed to find the id",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	profile, err := cr.userUseCase.ViewProfile(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "failed",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusAccepted, res.Response{
		StatusCode: 200,
		Message:    "Profile",
		Data:       profile,
		Errors:     nil,
	})
}
func (cr *UserHandler) EditProfile(c *gin.Context) {
	id, err := handlerUtil.GetUserIdFromContext(c)
	if err != nil {
		c.JSON(http.StatusAccepted, res.Response{
			StatusCode: 400,
			Message:    "failed to find the id",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	var update req.UserReq
	profile, err := cr.userUseCase.EditProfile(id, update)
	if err != nil {
		c.JSON(http.StatusBadRequest, res.Response{
			StatusCode: 400,
			Message:    "failed to update the profile",
			Data:       nil,
			Errors:     err.Error(),
		})
		return
	}
	c.JSON(http.StatusAccepted, res.Response{
		StatusCode: 200,
		Message:    "Updated successfully",
		Data:       profile,
		Errors:     nil,
	})
}
