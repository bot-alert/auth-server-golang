package controller

import (
	"auth-server-go/configuration"
	"auth-server-go/util"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"time"
)

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type Jwt struct {
	Token  string  `json:"token"`
	Expire float64 `json:"expire"`
}

func Auth(c *fiber.Ctx) error {
	var login Login

	err := json.Unmarshal(c.Body(), &login)
	if err != nil {
		return err
	}
	for _, user := range configuration.Users.Users {
		if user.Username == user.Username && user.Password == login.Password {
			token, err := util.GenerateJwtToken(user.Username)
			if err != nil {
				return err
			}
			jwt := Jwt{
				Token:  token,
				Expire: time.Hour.Seconds() * 2,
			}
			return c.Status(fiber.StatusOK).JSON(&jwt)
		}
	}
	return c.SendString("Hello, World!")
}
