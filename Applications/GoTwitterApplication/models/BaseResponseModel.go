package models

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type BaseResponseModel[T interface{}] struct {
	Data         T      `json:"data"`
	ErrorMessage string `json:"errorMessage"`
}

func BuildBadResponse(errorMessage string) BaseResponseModel[interface{}] {
	return BaseResponseModel[interface{}]{
		Data:         new(struct{}),
		ErrorMessage: errorMessage,
	}
}

func BuildGoodResponse[T interface{}](data T) BaseResponseModel[T] {
	return BaseResponseModel[T]{
		Data:         data,
		ErrorMessage: "",
	}
}

func SendGoodResponse[T interface{}](c echo.Context, data T) error {
	return c.JSON(http.StatusOK, BuildGoodResponse[T](data))
}

func SendBadResponse(c echo.Context, errorMessage string) error {
	return c.JSON(http.StatusBadRequest, BuildBadResponse(errorMessage))
}
