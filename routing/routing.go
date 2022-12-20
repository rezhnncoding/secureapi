package routing

import (
	"github.com/labstack/echo/v4"
	"puppy/controller"
)

func SetRouting(e *echo.Echo) error {

	RouteUserController(e)
	RouteAccountController(e)

	return nil
}

func RouteUserController(e *echo.Echo) {
	userController := controller.NewUserController()
	e.POST("/UploadPhoto", userController.UploadPhoto)
	g := e.Group("users")

	g.GET("/getList", userController.GetUserList)

	g.POST("/CreateUser", userController.CreateNewUser)

}

func RouteAccountController(e *echo.Echo) {
	accountController := controller.NewAccountController()
	e.POST("/login", accountController.LoginUser)
}
