package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wazven/backendblog/util"
)

func IsAuthenticate(c *fiber.Ctx) error{
	cookie := c.Cookies("jwt")

	if _, err :=util.ParseJwt(cookie); err != nil{
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message":"Unauthenticated",
		})
	}
	return c.Next()
}