package router

import "github.com/gofiber/fiber/v2"

// ParamFetcher is a function that fetches a parameter from the request
type ParamFetcher func(*fiber.Ctx) string

// FromPathParam returns a ParamFetcher that fetches a parameter from the path
func FromPathParam(paramName string) ParamFetcher {
	return func(c *fiber.Ctx) string {
		return c.Params(paramName)
	}
}
