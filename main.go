package main

import (
	"errors"
	"go_error_handling_playground/model"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
)

// errType query paramで受け取るエラータイプ
type errType string

const (
	errTypeBadRequest         errType = "bad_request"
	errTypeNotFound           errType = "not_found"
	errTypeInternalServer     errType = "internal_server"
	errTypeCarIsAlreadyBooked errType = "car_is_already_booked"
	errTypeShopClosed         errType = "shop_closed"
)

func NewErrType(t string) errType {
	return errType(t)
}

func main() {
	e := echo.New()
	// errorsパッケージを使ったサンプル
	e.GET("/errors", func(c echo.Context) error {
		slog.Info("Start GET /errors")
		t := c.QueryParam("type")
		err := handleErrByErrors(NewErrType(t))
		slog.Info("error check", "errType", t, "err", err)
		if err != nil {
			// 車を予約したいというビジネス要件を想定
			if errors.Is(err, model.ErrCarIsAlreadyBooked) {
				// 車が予約されているのに、再度予約しようとした場合は422を返す
				// その場合、FE側で予約ページにて、車が予約されている旨を表示する想定
				return c.JSON(http.StatusUnprocessableEntity, err)
			} else if errors.Is(err, model.ErrShopClosed) {
				// 店が閉まっている時は422を返す
				// その場合、FE側で予約ページにて、店が閉まっている旨を表示する想定
				return c.JSON(http.StatusUnprocessableEntity, err)
			}

			if errors.Is(err, model.ErrBadRequest) {
				return c.JSON(http.StatusBadRequest, err)
			} else if errors.Is(err, model.ErrNotFound) {
				return c.JSON(http.StatusNotFound, err)
			} else {
				return c.JSON(http.StatusInternalServerError, err)
			}
		}
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
