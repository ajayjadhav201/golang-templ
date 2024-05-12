package routes

import (
	"github.com/ajayjadhav201/golang-templ/handler"
	"github.com/labstack/echo/v4"
)

func RegisterRoutes(app *echo.Echo) {
	//
	GetRoutes(app)
	PostRoutes(app)
}

// GET Routes
func GetRoutes(app *echo.Echo) {
	userHandler := handler.UserHandler{}
	//
	app.GET("/", handler.WelcomePage)
	app.GET("/login", handler.LoginPage)
	app.GET("/signup", handler.SignupPage)
	app.GET("/contactus", handler.ContactUsPage)
	app.GET("/user", userHandler.HandleUserShow)
}

// POST rouets
func PostRoutes(app *echo.Echo) {
	app.POST("/login", handler.Login)
	app.GET("/darktheme", handler.ChangeTheme)
}
