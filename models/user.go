package models

import (
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"os"
)

type User struct {
	gorm.Model
	Id           uint   `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Username     string `json:"username" gorm:"unique"`
	PasswordHash string `json:"-"`
	DisplayName  string `json:"display_name"`
	RoleId       uint   `json:"role_id"`
	Role         Role   `json:"role" gorm:"foreignKey:RoleId"`
}

var (
	jwtKey = os.Getenv("JWT_KEY")
)

// HashPassword : Hash Password
func (u *User) HashPassword() {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(u.PasswordHash), bcrypt.DefaultCost)
	u.PasswordHash = string(bytes)
}

// GenerateToken : Generate Token
func (u *User) GenerateToken() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": u.Id,
	})

	tokenString, err := token.SignedString(jwtKey)
	return tokenString, err
}
