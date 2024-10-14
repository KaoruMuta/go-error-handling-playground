package main

import "go_error_handling_playground/model"

func handleErrByErrors(e errType) error {
	// 諸々のビジネスロジックが書かれており、返すエラーの種類が以下の想定
	switch e {
	case errTypeBadRequest:
		return model.AppErrBadRequest
	case errTypeNotFound:
		return model.AppErrNotFound
	case errTypeInternalServer:
		return model.AppErrInternalServer
	case errTypeCarIsAlreadyBooked:
		return model.AppErrCarIsAlreadyBooked
	case errTypeShopClosed:
		return model.AppErrShopClosed
	default:
		return model.AppErrInternalServer
	}
}
