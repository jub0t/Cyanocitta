package api

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo/v4"
)

// A Middle to parse JSON request BODY

func CustomBodyParser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		reqBodyRaw := c.Request().Body
		defer reqBodyRaw.Close()

		bodyBytes, err := ioutil.ReadAll(reqBodyRaw)
		if err != nil || len(bodyBytes) == 0 {
			return next(c)
		}

		var parsedBody map[string]interface{}
		if err := json.Unmarshal(bodyBytes, &parsedBody); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, errors.New("invalid JSON format"))
		}

		c.Set("Body", parsedBody)

		return next(c)
	}
}
