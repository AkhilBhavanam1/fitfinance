package fitness

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func FitnessHome(ctx echo.Context) error {

	return ctx.String(http.StatusOK, "This is fitness home")
}
