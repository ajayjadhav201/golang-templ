package handler

import (
	"github.com/ajayjadhav201/golang-templ/model"
	"github.com/ajayjadhav201/golang-templ/view/user"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
}

func (h UserHandler) HandleUserShow(c echo.Context) error {
	u := model.User{
		Email: "ajayjadhav201@gmail.com",
	}
	return render(c, user.Show(u))
}
