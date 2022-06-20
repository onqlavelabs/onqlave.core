package utils

import "github.com/labstack/echo/v4"

type HandlerSetup interface {
	SetupRoutes(groupName string, group *echo.Group) []Subscription
}
