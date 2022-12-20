package Utility

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {

	validate := validator.New()
	errvalidate := validate.Struct(i)
	if errvalidate != nil {
		if strings.Contains(errvalidate.Error(), "FirstName") {
			return echo.NewHTTPError(http.StatusBadRequest, "firstname is not valid")
		} else if strings.Contains(errvalidate.Error(), "LastName") {
			return echo.NewHTTPError(http.StatusBadRequest, "LastName is not valid")
		} else if strings.Contains(errvalidate.Error(), "BirthDate") {
			return echo.NewHTTPError(http.StatusBadRequest, "BirthDate is not a valid birth number")
		} else if strings.Contains(errvalidate.Error(), "Code") {
			return echo.NewHTTPError(400, "code melli is not valid please enter a valid format")
		} else if strings.Contains(errvalidate.Error(), "UserName") {
			return echo.NewHTTPError(400, "username is not valid or used before please try again")
		} else if strings.Contains(errvalidate.Error(), "Password") {
			return echo.NewHTTPError(400, "Password is not correct or have special carcters please try again")
		}

		//return echo.NewHTTPError(http.StatusBadRequest, errvalidate.Error())
	}

	return nil
}
