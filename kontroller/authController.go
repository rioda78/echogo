package kontroller

import (
	"echogo/models"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func Register(c echo.Context) error {
	type RequestBody struct {
		Username  string `json:"username" validate:"required"`
		Password  string `json:"password" validate:"required"`
		Firstname string `json:"first_name"`
		Lastname  string `json:"last_name"`
		Peran     uint   `json:"-"`
	}
	var body RequestBody

	if err := c.Bind(&body); err != nil {
		return err
	}
	if err := c.Validate(&body); err != nil {
		return err
	}

	db, _ := c.Get("db").(*gorm.DB)

	if err := db.Where("username = ?", body.Username).First(&models.User{}).Error; err == nil {
		return c.NoContent(http.StatusConflict)
	}
	user := models.User{
		Username:     body.Username,
		PasswordHash: body.Password,
		FirstName:    body.Firstname,
		LastName:     body.Lastname,
		RoleId:       3,
	}

	user.HashPassword()
	db.Create(&user)

	token, _ := user.GenerateToken()

	var cookie http.Cookie

	cookie.Name = "token"
	cookie.Value = token
	cookie.Expires = time.Now().Add(7 * 24 * time.Hour)

	c.SetCookie(&cookie)

	return c.JSON(http.StatusOK, echo.Map{
		"token": token,
		"user":  user,
	})
}

// Login : Login Router
func Login(c echo.Context) error {
	type RequestBody struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`
	}

	var body RequestBody

	if err := c.Bind(&body); err != nil {
		return c.NoContent(http.StatusNotFound)
	}
	if err := c.Validate(&body); err != nil {
		return err
	}

	db, _ := c.Get("db").(*gorm.DB)

	var user models.User

	if err := db.Where("username = ?", body.Username).First(&user).Error; err != nil {
		return c.NoContent(http.StatusConflict)
	}

	if bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(body.Password)) != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	token, _ := user.GenerateToken()

	var cookie http.Cookie

	cookie.Name = "token"
	cookie.Value = token
	cookie.Expires = time.Now().Add(7 * 24 * time.Hour)

	c.SetCookie(&cookie)

	return c.JSON(http.StatusOK, echo.Map{
		"token": token,
		"user":  user,
	})
}

// Logout : Logout Router
func Logout(c echo.Context) error {
	tokenCookie, _ := c.Get("tokenCookie").(*http.Cookie)

	tokenCookie.Value = ""
	tokenCookie.Expires = time.Unix(0, 0)

	c.SetCookie(tokenCookie)

	return c.NoContent(200)
}
