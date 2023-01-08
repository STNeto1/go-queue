package auth

import (
	lib "_lib"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type LoginRequestBody struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (router AuthRouter) LoginHandler(c *gin.Context) {
	body := LoginRequestBody{}

	if err := c.ShouldBindJSON(&body); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]lib.ApiError, len(ve))
			for i, fe := range ve {
				out[i] = lib.ApiError{Param: fe.Field(), Message: lib.MsgForTag(fe)}
			}

			c.JSON(http.StatusBadRequest, lib.NewBadValidation(out))
			return
		}
	}

	usr, err := router.as.LoginUser(c.Request.Context(), body.Email, body.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, lib.BadRequest{
			Message:    "Invalid Credentials",
			StatusCode: http.StatusBadRequest,
		})
		return
	}

	token, err := router.as.GenerateToken(usr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, lib.InternalServerError{
			Message:    "Internal Server Error",
			StatusCode: http.StatusInternalServerError,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"token": token,
	})
}
