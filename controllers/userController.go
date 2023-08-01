package controllers

import (
	"echo-crud/initializers"
	"echo-crud/models"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/teris-io/shortid"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"time"
)

func CreateUser(c echo.Context) error {
	user := models.User{}

	signupInfo := models.UserInfo{}

	if err := c.Bind(&signupInfo); err != nil {
		return err
	}

	if signupInfo.Name == "" {
		return c.String(http.StatusBadRequest, "名前が入力されていません。")
	}
	user.Name = signupInfo.Name
	if signupInfo.Email == "" {
		return c.String(http.StatusBadRequest, "メールアドレスが入力されていません。")
	}
	user.Email = signupInfo.Email

	// create short id
	sid, _ := shortid.New(1, shortid.DefaultABC, 2342)
	user.ID, _ = sid.Generate()

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		return err
	}

	user.Password = string(hash)

	initializers.DB.Create(&user)
	return c.JSON(http.StatusCreated, user)
}

func LoginUser(c echo.Context) error {
	var user models.User
	var loginInfo models.UserInfo

	if err := c.Bind(&loginInfo); err != nil {
		return err
	}

	initializers.DB.First(&user, "email = ?", loginInfo.Email)

	if user.ID == "" {
		return c.String(http.StatusBadRequest, "メールアドレスもしくはパスワードが間違っています。")
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginInfo.Password))

	if err != nil {
		return c.String(http.StatusBadRequest, "メールアドレスもしくはパスワードが間違っています。")
	}

	// create token

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = user.ID
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	//// エンコードされたトークンを生成し、応答として送信します。
	//// 署名文字列は秘密である必要があります (生成された UUID も機能します)
	t, err := token.SignedString([]byte(os.Getenv("secret")))
	if err != nil {
		return err
	}
	if t != "" {
		return c.JSON(http.StatusOK, echo.Map{
			"token": t,
		})
	}
	return echo.ErrUnauthorized

}
