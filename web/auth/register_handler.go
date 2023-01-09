package auth

import (
	"_core/auth"
	lib "_lib"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/labstack/echo/v4"
)

type RegisterRequestBody struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (router AuthRouter) RegisterHandler(c echo.Context) error {
	body := new(RegisterRequestBody)

	if err := c.Bind(body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, lib.BadRequest{
			Message:    "Invalid request body",
			StatusCode: http.StatusBadRequest,
		})
	}

	if err := c.Validate(body); err != nil {
		errors := lib.ParseValidatorErrors(err)
		return echo.NewHTTPError(http.StatusBadRequest, lib.BadValidation{
			Message:    "Invalid request body",
			StatusCode: http.StatusBadRequest,
			Errors:     errors,
		})
	}

	usr, err := router.as.RegisterUser(c.Request().Context(), auth.RegisterUserPayload{
		Name:     body.Name,
		Email:    body.Email,
		Password: body.Password,
	})
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
