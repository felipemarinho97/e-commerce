package db

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/felipemarinho97/e-commerce/common"
	"github.com/felipemarinho97/e-commerce/config"
	"go.uber.org/atomic"
)

// Database is the interface for the database.
type Database interface {
	GetProduct(ctx context.Context, id, quantity int32) (ProductResponse, error)
	GetGift(ctx context.Context) (ProductResponse, error)
}

// DB is the database implementation.
type DB struct {
	Products *map[int32]*Product
	Gifts    *map[int32]*Product
}

// Product is the product model.
type Product struct {
	ID          int32         `json:"id"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Amount      atomic.Uint32 `json:"amount"`
	Price       int32         `json:"price"`
	IsGift      bool          `json:"is_gift"`
}

// ProductResponse is the response for the GetProduct method.
type ProductResponse struct {
	ID     int32 `json:"id"`
	Amount int32 `json:"amount"`
	Price  int32 `json:"price"`
}

// New creates a new database reading the products from the given file.
func New() (Database, error) {
	out, err := ioutil.ReadFile(config.Get().ProductsMockFile)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %s", err.Error())
	}
	var products []Product
	err = json.Unmarshal(out, &products)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling file: %s", err.Error())
	}

	var pMap map[int32]*Product = make(map[int32]*Product, len(products))
	var giftMap map[int32]*Product = make(map[int32]*Product)

	for _, product := range products {
		product := product
		pMap[product.ID] = &product
		if product.IsGift {
			giftMap[product.ID] = &product
		}
	}

	return &DB{
		Products: &pMap,
		Gifts:    &giftMap,
	}, nil
}

// GetProduct returns the product with the given id.
func (db *DB) GetProduct(ctx context.Context, id, quantity int32) (ProductResponse, error) {
	product, ok := (*db.Products)[id]
	if !ok {
		return ProductResponse{}, common.ErrProductNotFound
	}
	if product.IsGift {
		return ProductResponse{}, common.ErrProductIsGift
	}
	if product.Amount.Load() == 0 {
		return ProductResponse{}, common.ErrProductOutOfStock
	}
	p := ProductResponse{
		ID:    product.ID,
		Price: product.Price,
	}

	if amount := product.Amount.Load(); amount < uint32(quantity) {
		product.Amount.CAS(amount, 0)
		p.Amount = int32(amount)
	} else {
		product.Amount.Sub(uint32(quantity))
		p.Amount = quantity
	}

	return p, nil
}

// GetGift returns the next available gift product.
func (db *DB) GetGift(ctx context.Context) (ProductResponse, error) {
	for _, product := range *db.Gifts {
		if product.Amount.Load() > 0 {
			product.Amount.Sub(1)
			return ProductResponse{
				ID:     product.ID,
				Amount: 1,
				Price:  product.Price,
			}, nil
		}
	}
	return ProductResponse{}, common.ErrGiftsOutOfStock
}
