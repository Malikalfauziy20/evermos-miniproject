package middleware

import "github.com/gofiber/fiber/v2"

func AdminOnly() fiber.Handler {
	return func(c *fiber.Ctx) error {
		isAdmin := c.Locals("is_admin").(bool)
		if !isAdmin {
			return c.Status(403).JSON(fiber.Map{"status": false, "message": "hanya admin yang boleh mengakses"})
		}
		return c.Next()
	}
}
