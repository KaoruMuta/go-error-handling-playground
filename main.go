package main

import (
	"errors"
	"go_error_handling_playground/model"
	"log/slog"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/morikuni/failure/v2"
)

// errType query paramで受け取るエラータイプ
type errType string

func (t errType) String() string {
	return string(t)
}

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

func newStatusCode(err error) int {
	switch failure.CodeOf(err) {
	case model.ErrCodeBadRequest:
		return http.StatusBadRequest
	case model.ErrCodeNotFound:
		return http.StatusNotFound
	case model.ErrCodeCarIsAlreadyBooked, model.ErrCodeShopClosed:
		return http.StatusUnprocessableEntity
	case model.ErrCodeInternalServer:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
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
	// failureパッケージを使ったサンプル
	e.GET("/failure", func(c echo.Context) error {
		slog.Info("Start GET /failure")
		t := c.QueryParam("type")
		err := handleErrByFailure(NewErrType(t))
		slog.Info("error check", "errType", t, "err", err)
		if err != nil {
			// failureを使うと
			// - エラーの呼び出し元の情報を出力できる
			// - errors#Isの比較のためにカスタムエラーを逐一定義する必要がない (エラーコード(識別子)とステータスコードのマッピングは必要)
			// - 識別子やメッセージ、パラメータを取得できるので、それらを元にエラーレスポンスを生成できる
			//
			// ```bash
			// # fmt.Printf(%+v, err)の出力例
			// 2024/10/14 16:19:43 INFO Start GET /failure
			// [main.handleErrByFailure] /Users/k_muta/Desktop/PersonalProject/go_error_handling_playground/failure.go:23
			//     Unknown error
			//     {errType=}
			// model.appError("Internal server error")
			// [CallStack]
			//     [main.handleErrByFailure] /Users/k_muta/Desktop/PersonalProject/go_error_handling_playground/failure.go:23
			//     [main.main.func2] /Users/k_muta/Desktop/PersonalProject/go_error_handling_playground/main.go:82
			//     [v4.(*Echo).add.func1] /Users/k_muta/go/pkg/mod/github.com/labstack/echo/v4@v4.12.0/echo.go:587
			//     [v4.(*Echo).ServeHTTP] /Users/k_muta/go/pkg/mod/github.com/labstack/echo/v4@v4.12.0/echo.go:674
			//     [http.serverHandler.ServeHTTP] /usr/local/go/src/net/http/server.go:3210
			//     [http.(*conn).serve] /usr/local/go/src/net/http/server.go:2092
			//     [runtime.goexit] /usr/local/go/src/runtime/asm_arm64.s:1223
			// ```
			if code, ok := failure.CodeOf(err).(model.ErrorCode); ok {
				// - status_codeをfailureで生成したエラーから割り出す (エラーコードとstatus_codeのマッピング)
				// - エラーレスポンスのモデルをfailureで生成したエラーから割り出す
				return c.JSON(newStatusCode(err), model.NewAppError(code, string(failure.MessageOf(err))))
			}
			return c.JSON(http.StatusInternalServerError, "unknown error code")
		}
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
