package main

import "go_error_handling_playground/model"

func handleErrByErrors(e errType) error {
	// 諸々のビジネスロジックが書かれており、返すエラーの種類が以下の想定
	switch e {
	case errTypeBadRequest:
		return model.ErrBadRequest
	case errTypeNotFound:
		return model.ErrNotFound
	case errTypeInternalServer:
		return model.ErrInternalServer
	case errTypeCarIsAlreadyBooked:
		return model.ErrCarIsAlreadyBooked
	case errTypeShopClosed:
		return model.ErrShopClosed
	default:
		return model.ErrInternalServer
	}
}
