package main

import (
	"go_error_handling_playground/model"

	"github.com/morikuni/failure/v2"
)

func handleErrByFailure(e errType) error {
	// 諸々のビジネスロジックが書かれており、返すエラーの種類が以下の想定
	switch e {
	case errTypeBadRequest:
		return failure.New(model.ErrCodeBadRequest, failure.Message("Bad request"))
	case errTypeNotFound:
		return failure.New(model.ErrCodeNotFound, failure.Message("Not found"))
	case errTypeInternalServer:
		return failure.New(model.ErrCodeInternalServer, failure.Message("Internal server error"))
	case errTypeCarIsAlreadyBooked:
		return failure.New(model.ErrCodeCarIsAlreadyBooked, failure.Message("Car is already booked"))
	case errTypeShopClosed:
		return failure.New(model.ErrCodeShopClosed, failure.Message("Shop is closed"))
	default:
		return failure.New(model.ErrCodeInternalServer, failure.Message("Unknown error"), failure.Context{"errType": e.String()})
	}
}
