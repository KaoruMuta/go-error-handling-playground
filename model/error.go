package model

// アプリケーション側で起こりうる特定のエラーに対して特別ハンドリングしたい場合、独自エラーとして定義しておく
var (
	AppErrBadRequest         = NewAppError(ErrCodeBadRequest, "Bad request")
	AppErrNotFound           = NewAppError(ErrCodeNotFound, "Not found")
	AppErrInternalServer     = NewAppError(ErrCodeInternalServer, "Internal server error")
	AppErrCarIsAlreadyBooked = NewAppError(ErrCodeCarIsAlreadyBooked, "Car is already booked")
	AppErrShopClosed         = NewAppError(ErrCodeShopClosed, "Shop is closed")
)

type ErrorCode string

const (
	ErrCodeBadRequest         ErrorCode = "bad_request"
	ErrCodeNotFound           ErrorCode = "not_found"
	ErrCodeInternalServer     ErrorCode = "internal_server"
	ErrCodeCarIsAlreadyBooked ErrorCode = "car_is_already_booked"
	ErrCodeShopClosed         ErrorCode = "shop_closed"
)

type appError struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
}

// Error errors#Isで比較できるように、appErrorがerrorインターフェースを満たすようにする
func (e appError) Error() string {
	return e.Message
}

func NewAppError(code ErrorCode, message string) appError {
	return appError{
		Code:    code,
		Message: message,
	}
}
