package auth

import (
	lib "_lib"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
)

type LoginRequestBody struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (router AuthRouter) LoginHandler(c echo.Context) error {
	body := new(LoginRequestBody)

	if err := c.Bind(body); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, lib.BadRequest{
			Message:    "Invalid request body",
			StatusCode: http.StatusBadRequest,
		})
	}
	if err := c.Validate(body); err != nil {
		log.Println(err)
		errors := lib.ParseValidatorErrors(err)
		return echo.NewHTTPError(http.StatusBadRequest, lib.BadValidation{
			Message:    "Invalid request body",
			StatusCode: http.StatusBadRequest,
			Errors:     errors,
		})
	}

	usr, err := router.as.LoginUser(c.Request().Context(), body.Email, body.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, lib.BadRequest{
			Message:    "Invalid Credentials",
			StatusCode: http.StatusBadRequest,
		})
		return err
	}

	token, err := router.as.GenerateToken(usr)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, lib.InternalServerError{
			Message:    "Internal Server Error",
			StatusCode: http.StatusInternalServerError,
		})
	}

	return c.JSON(http.StatusCreated, gin.H{
		"token": token,
	})
}
