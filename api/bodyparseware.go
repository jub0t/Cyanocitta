package api

import (
	"disco/structs"
	"encoding/json"
	"io/ioutil"

	"github.com/labstack/echo/v4"
)

// A Middle to parse JSON request BODY

func CustomBodyParser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		reqBodyRaw := c.Request().Body
		defer reqBodyRaw.Close()

		bodyBytes, err := ioutil.ReadAll(reqBodyRaw)
		if err != nil || len(bodyBytes) == 0 {
			return c.JSON(200, structs.Response{
				Success: false,
				Message: "Body required but not found",
			})
		}

		var parsedBody structs.AnyData
		if err := json.Unmarshal(bodyBytes, &parsedBody); err != nil {
			return c.JSON(200, structs.Response{
				Success: false,
				Message: "Invalid Body",
			})

			// echo.NewHTTPError(http.StatusBadRequest, errors.New("invalid JSON format"))
		}

		c.Set("Body", parsedBody)

		return next(c)
	}
}
