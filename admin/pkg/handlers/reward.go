package handler

import "github.com/labstack/echo/v4"

func NewReward(c echo.Context) error {
	return c.JSON(200, "NewReward reward")
}

func GetReward(c echo.Context) error {
	return c.JSON(200, "Get reward")
}

func UpdateReward(c echo.Context) error {
	return c.JSON(200, "UpdateReward reward")
}

func DeleteReward(c echo.Context) error {
	return c.JSON(200, "DeleteReward reward")
}
