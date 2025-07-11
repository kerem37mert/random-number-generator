package handlers

import (
	"math/rand/v2"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Value int `json:"value"`
}

type Error struct {
	Message string `json:"message"`
}

func RandomNumber(c echo.Context) error {

	minStr := c.QueryParam("min")
	maxStr := c.QueryParam("max")

	min, minErr := strconv.Atoi(minStr)
	max, maxErr := strconv.Atoi(maxStr)
	if minErr != nil || maxErr != nil {
		return c.JSON(http.StatusBadRequest, &Error{"invalid parameters"})
	}

	randomNumber := rand.IntN(max-min+1) + min

	response := &Response{Value: randomNumber}
	return c.JSON(http.StatusOK, response)
}
