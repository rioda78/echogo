package kontroller

import (
	"echogo/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func Register(c echo.Context) error {
	type RequestBody struct {
		Username string `json:"username" validate:"required"`
		Password string `json:"password" validate:"required"`

		DisplayName string `json:"display_name" validate:"required"`
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
		DisplayName:  body.DisplayName,
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
