package api

import (
	database "disco/db"
	"disco/structs"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func TokenMiddleware(db *gorm.DB) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		headers := ctx.Request().Header()
		token := headers["Token"]

		if len(token) <= 0 {
			return ctx.JSON(structs.Response{
				Success: false,
				Message: "Invalid Token, Unauthenticated",
			})
		}

		if user, err := database.GetUserByToken(db, token); err != nil {
			return ctx.JSON(structs.Response{
				Success: false,
				Message: "Internal Server Error",
			})
		} else {
			if (user != structs.User{}) {
				ctx.Locals("User", user)
			} else {
				return ctx.JSON(structs.Response{
					Success: false,
					Message: "Unauthenticated",
				})
			}
		}

		return ctx.Next()
	}
}
