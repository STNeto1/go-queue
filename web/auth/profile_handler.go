package auth

import (
	lib "_lib"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (router AuthRouter) ProfileHandler(c echo.Context) error {
	usr, err := router.UserFromContext(c)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, lib.Unauthorized{
			Message:    "Unauthorized",
			StatusCode: http.StatusUnauthorized,
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"id":         usr.ID.String(),
		"name":       usr.Name,
		"email":      usr.Email,
		"created_at": usr.CreatedAt.String(),
	})
}
