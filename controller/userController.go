package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"puppy/Utility"
	httpResponse "puppy/models/common/httpresponse"
	userViewModel "puppy/models/user"
	"puppy/service"
)

type UserController interface {
	GetUserList(c echo.Context) error
	CreateNewUser(c echo.Context) error
	UploadPhoto(c echo.Context) error
}

type userController struct {
}

func NewUserController() UserController {
	return userController{}
}

func (uc userController) GetUserList(c echo.Context) error {
	apiContext := c.(*Utility.ApiContext)
	fmt.Println(apiContext.GetUserId())

	userService := service.NewUserService()
	userList, err := userService.GetUserList()
	if err != nil {
		println(err)
	}

	return c.JSON(http.StatusOK, httpResponse.SuccessResponse(userList))
}
func (uc userController) CreateNewUser(c echo.Context) error {
	apiContext := c.(*Utility.ApiContext)
	operatorUserId, err := apiContext.GetUserId()
	if err != nil {
		return c.JSON(http.StatusBadRequest, "")
	}
	userService := service.NewUserService()
	//isValid := userService.IsUserValidForAccess(operatorUserId, "CreateUser")
	//if !isValid {
	//	return c.JSON(http.StatusForbidden, "")
	//}

	newUser := new(userViewModel.CreateNewUserViewModel)

	if err := c.Bind(newUser); err != nil {
		return c.JSON(http.StatusBadRequest, httpResponse.SuccessResponse("Data not found"))
	}

	if err := c.Validate(newUser); err != nil {
		return c.JSON(http.StatusBadRequest, httpResponse.SuccessResponse(err))
	}

	file, err := apiContext.FormFile("file")
	if err == nil {
		src, err := file.Open()
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		wd, err := os.Getwd()
		imageServerPath := filepath.Join(wd, "wwwroot", "images", "userAvatar", "dcfvgbhnjm"+filepath.Ext(file.Filename))

		des, err := os.Create(imageServerPath)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		defer des.Close()

		_, err = io.Copy(des, src)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		newUser.AvatarName = "dcfvgbhnjm" + filepath.Ext(file.Filename)
	}

	newUser.CreatorUserId = operatorUserId

	newUserId, err := userService.CreateNewUser(*newUser)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	userResData := struct {
		NewUserId string
	}{
		NewUserId: newUserId,
	}
	return c.JSON(http.StatusOK, httpResponse.SuccessResponse(userResData))
}

func (uc userController) UploadPhoto(c echo.Context) error {
	apiContext := c.(*Utility.ApiContext)

	file, err := apiContext.FormFile("file")
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	des, err := os.Create(file.Filename)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	defer des.Close()

	_, err = io.Copy(des, src)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	userResData := struct {
		IsSuccess bool
	}{
		IsSuccess: true,
	}

	return c.JSON(http.StatusOK, userResData)
}
