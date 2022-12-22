package http

import (
	"context"
	"net/http"
	"strconv"

	"github.com/alramdein/user-service/model"
	"github.com/alramdein/user-service/usecase"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
)

type UserHandler struct {
	UserUsecase model.UserUsecase
}

func NewUserHandler(e *echo.Echo, u model.UserUsecase) {
	handler := &UserHandler{
		UserUsecase: u,
	}

	g := e.Group("/users")
	g.POST("/:user_id", handler.FindByID)
	g.POST("/create", handler.Create)
	g.POST("/update", handler.Update)
	g.POST("/delete", handler.Delete)
}

func (u *UserHandler) FindByID(c echo.Context) error {
	userID, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusBadRequest, &Response{
			Message: "invalid userID",
		})
	}

	user, err := u.UserUsecase.FindByID(context.Background(), userID)
	switch err {
	case nil:
		return c.JSON(http.StatusOK, user)
	case usecase.ErrNotFound:
		return c.JSON(http.StatusNotFound, &Response{
			Message: "user not found",
		})
	default:
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, &Response{
			Message: "something went wrong",
		})
	}
}

func (u *UserHandler) Create(c echo.Context) error {
	username := c.FormValue("username")
	email := c.FormValue("email")
	password := c.FormValue("password")
	roleID, err := strconv.ParseInt(c.FormValue("role_id"), 10, 64)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusBadRequest, &Response{
			Message: "invalid roleID",
			Data:    nil,
		})
	}

	err = u.UserUsecase.Create(context.Background(), model.CreateUserInput{
		Username: username,
		Email:    email,
		Password: password,
		RoleID:   roleID,
	})
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, &Response{
			Message: "failed to create user",
		})
	}

	return c.JSON(http.StatusOK, &Response{
		Message: "successfully create user",
	})
}

func (u *UserHandler) Update(c echo.Context) error {
	username := c.FormValue("username")
	email := c.FormValue("email")
	password := c.FormValue("password")
	userID, err := strconv.ParseInt(c.FormValue("user_id"), 10, 64)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusBadRequest, &Response{
			Message: "invalid userID",
			Data:    nil,
		})
	}
	roleID, err := strconv.ParseInt(c.FormValue("role_id"), 10, 64)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusBadRequest, &Response{
			Message: "invalid roleID",
			Data:    nil,
		})
	}

	err = u.UserUsecase.Update(context.Background(), model.User{
		ID:       userID,
		Username: username,
		Email:    email,
		Password: password,
		RoleID:   roleID,
	})
	switch err {
	case nil:
		return c.JSON(http.StatusOK, &Response{
			Message: "successfully update user",
		})
	case usecase.ErrUserNotFound:
		return c.JSON(http.StatusNotFound, &Response{
			Message: "user not found",
		})
	default:
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, &Response{
			Message: "failed to update user",
		})
	}
}

func (u *UserHandler) Delete(c echo.Context) error {
	userID, err := strconv.ParseInt(c.FormValue("user_id"), 10, 64)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusBadRequest, &Response{
			Message: "invalid userID",
			Data:    nil,
		})
	}

	err = u.UserUsecase.Delete(context.Background(), userID)
	switch err {
	case nil:
		return c.JSON(http.StatusOK, &Response{
			Message: "successfully delete user",
		})
	case usecase.ErrUserNotFound:
		return c.JSON(http.StatusNotFound, &Response{
			Message: "user not found",
		})
	default:
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, &Response{
			Message: "failed to delete user",
		})
	}
}
