package handler

import (
	"github.com/ajayjadhav201/golang-templ/common"
	"github.com/ajayjadhav201/golang-templ/view/home"
	"github.com/labstack/echo/v4"
)

func WelcomePage(c echo.Context) error {
	theme := common.GetCookie(c, "theme", "light") == "dark"
	common.Println("Rednering base page: ", theme)
	return render(c, home.WelcomePage())
}

func LoginPage(c echo.Context) error {
	return render(c, home.LoginPage())
}
func SignupPage(c echo.Context) error {
	return render(c, home.SignupPage())
}
func ContactUsPage(c echo.Context) error {
	return render(c, home.ContactUs())
}

//
//
//
//
//
//
//
//
//

func Login(c echo.Context) error {
	//
	username := c.FormValue("username")
	password := c.FormValue("password")
	//
	// time.Sleep(2 * time.Second)
	common.Println("ajaj username: ", username, "and password: ", password)
	return LoginPage(c)
}

func ChangeTheme(c echo.Context) error {
	theme := common.GetCookie(c, "theme", "light")
	if theme == "dark" {
		theme = "light"
	} else {
		theme = "dark"
	}
	c.SetCookie(common.Cookie("theme", theme))

	return WelcomePage(c)
}

// go build -o ./tmp/main.exe cmd/main.go
