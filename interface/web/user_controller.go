package web

import (
	"net/http"

	"github.com/mkaiho/go-oauth-sandbox/usecase"
)

type UserController struct {
	userUsecase usecase.UserUsecase
}

func NewUserController() Controller {
	return &UserController{userUsecase: usecase.NewUserUsecase()}
}

func (c *UserController) RegisterRoutes(router Router) {
	router.Get("users/:id", c.Find)
}

func (c *UserController) Find(context Contexter) {
	id := context.Param("id")
	user := c.userUsecase.Find(id)

	context.JSON(http.StatusOK, map[string]string{"id": user.ID()})
}
