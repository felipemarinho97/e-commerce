package db

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/felipemarinho97/e-commerce/common"
	"github.com/felipemarinho97/e-commerce/config"
	"go.uber.org/atomic"
)

type Database interface {
	GetProduct(ctx context.Context, id, quantity int64) (ProductResponse, error)
}

type DB struct {
	products *map[int64]*Product
}

type Product struct {
	ID          int64         `json:"id"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Amount      atomic.Uint64 `json:"amount"`
	Price       int64         `json:"price"`
	IsGift      bool          `json:"is_gift"`
}

type ProductResponse struct {
	ID     int64 `json:"id"`
	Amount int64 `json:"amount"`
	Price  int64 `json:"price"`
}

func New() Database {
	out, err := ioutil.ReadFile(config.Get().ProductsMockFile)
	if err != nil {
		common.Logger.LogFatal("error reading file", err.Error())
		os.Exit(-1)
	}
	var products []Product
	err = json.Unmarshal(out, &products)
	if err != nil {
		common.Logger.LogFatal("error unmarshalling file", err.Error())
		os.Exit(-1)
	}

	var pMap map[int64]*Product = make(map[int64]*Product, len(products))

	for _, product := range products {
		product := product
		pMap[product.ID] = &product
	}

	return &DB{
		products: &pMap,
	}
}

func (db *DB) GetProduct(ctx context.Context, id, quantity int64) (ProductResponse, error) {
	product, ok := (*db.products)[id]
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

	if amount := product.Amount.Load(); amount < uint64(quantity) {
		product.Amount.CAS(amount, 0)
		p.Amount = int64(amount)
	} else {
		product.Amount.Sub(uint64(quantity))
		p.Amount = quantity
	}

	return p, nil
}
