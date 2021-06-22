package views

import (
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"
)

// BindNoId wraps c.Bind(), but add the indication that Id is not needed
func BindNoId(c echo.Context, params interface{}) error {
	return c.Bind(params)
}

// QueryParamBool checks if `name` is a non empty query param
func QueryParamBool(c echo.Context, name string) bool {
	return c.QueryParam(name) != ""
}

// QueryParamInt64 parse the query param `name` to an int64
func QueryParamInt64(c echo.Context, name string) (int64, error) {
	idS := c.QueryParam(name)
	id, err := strconv.ParseInt(idS, 10, 64)
	if err != nil {
		return 0, fmt.Errorf("Impossible de décrypter l'ID reçu %s : %s", idS, err)
	}
	return id, nil
}
