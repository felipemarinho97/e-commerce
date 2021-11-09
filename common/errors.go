package common

import "errors"

var (
	ErrProductNotFound   = errors.New("product not found")
	ErrProductIsGift     = errors.New("product is gift")
	ErrProductOutOfStock = errors.New("product is out of stock")
	ErrGiftsOutOfStock   = errors.New("gifts are out of stock")
)
