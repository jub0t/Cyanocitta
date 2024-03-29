package api

import (
	database "disco/db"
	"disco/structs"
	"disco/utils"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type AuthBody struct {
	Username string
	Password string
}

func LoginRoute(db *gorm.DB) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		body := ctx.Get("Body").(structs.AnyData)
		username := body["Username"].(string)

		if len(username) < 3 {
			ctx.JSON(200, structs.Response{
				Success: false,
				Message: "Username must be atleast 3 characters",
			})
		}

		if user, err := database.GetUserByUsername(db, username); err != nil {
			return ctx.JSON(200, structs.Response{
				Success: false,
				Message: "Unable to fetch user info",
			})
		} else {
			if (user != structs.User{}) {
				return ctx.JSON(200, structs.Response{
					Success: false,
					Data: structs.AnyData{
						"Token": user.Token,
					},
				})
			} else {
				return ctx.JSON(200, structs.Response{
					Success: false,
					Message: "Unable to fetch user info",
				})
			}
		}
	}
}

func RegisterRoute(db *gorm.DB) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		body := ctx.Get("Body").(structs.AnyData)
		username := body["Username"].(string)

		if len(username) < 3 {
			ctx.JSON(200, structs.Response{
				Success: false,
				Message: "Username must be atleast 3 characters",
			})
		}

		if database.UserExists(db, username) {
			return ctx.JSON(200, structs.Response{
				Success: false,
				Message: "Username is taken",
			})
		}

		password := body["Password"].(string)

		if !utils.VerifyPass(password) {
			// Password Invalid
			return ctx.JSON(200, structs.Response{
				Success: false,
				Message: "Password must match criteria",
			})
		}

		token := utils.RandomString(128)
		db.Create(&structs.User{
			Username: username,
			Password: password,
			Token:    token,
		})

		return ctx.JSON(200, structs.Response{
			Success: true,
			Message: "Registered new account",
			Data: structs.AnyData{
				"Token": token,
			},
		})
	}
}
