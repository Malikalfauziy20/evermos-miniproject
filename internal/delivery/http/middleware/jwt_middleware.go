package middleware

import (
	"strings"

	appjwt "github.com/Malikalfauziy20/evermos-miniproject/pkg/jwt"
	"github.com/gofiber/fiber/v2"
)

func JWTProtected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status": false, "message": "token tidak ditemukan",
			})
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := appjwt.ParseToken(tokenStr)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status": false, "message": "token tidak valid atau kadaluarsa",
			})
		}

		userIDFloat, ok := claims["user_id"].(float64)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"status": false, "message": "token tidak valid",
			})
		}
		isAdmin, _ := claims["is_admin"].(bool)

		c.Locals("user_id", uint(userIDFloat))
		c.Locals("is_admin", isAdmin)

		return c.Next()
	}
}
