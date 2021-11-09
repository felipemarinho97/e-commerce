package api

import (
	"context"
	"errors"
	"fmt"
	"math"

	"github.com/felipemarinho97/e-commerce/common"
	"github.com/felipemarinho97/e-commerce/db"
	"github.com/felipemarinho97/e-commerce/discount"
	pb "github.com/felipemarinho97/e-commerce/examples/go/protos/api"
)

func calculateDiscount(ctx context.Context, productID, price int64) int64 {
	d, err := discount.GetDiscountPercentage(ctx, int32(productID))
	if err != nil {
		common.Logger.LogError("calculateDiscount", "Error getting discount percentage", err.Error())
		return 0
	}

	return int64(math.Ceil(float64(price) * float64(d)))
}

func (s server) Checkout(ctx context.Context, in *pb.CheckoutRequest) (*pb.CheckoutResponse, error) {
	response := &pb.CheckoutResponse{}
	for _, item := range in.Products {
		product, err := s.db.GetProduct(ctx, item.Id, item.Quantity)
		if errors.Is(err, common.ErrProductOutOfStock) {
			product = db.ProductResponse{
				ID:     item.Id,
				Amount: -1,
			}
			common.Logger.LogInfo("Checkout", "unable to check out product", err.Error())
		} else if err != nil {
			common.Logger.LogError("Checkout", "error checking out product", err.Error())
			common.Logger.LogDebug("Checkout", fmt.Sprintf("id=%d", item.Id), fmt.Sprintf("quantity=%d", item.Quantity))
			return nil, fmt.Errorf("error checking out product with id=%d: %s", item.Id, err.Error())
		}

		p := &pb.ProductResponse{
			Id:          product.ID,
			Quantity:    product.Amount,
			UnitAmount:  product.Price,
			Discount:    calculateDiscount(ctx, product.ID, product.Price),
			TotalAmount: product.Price * product.Amount,
			IsGift:      false,
		}
		response.Products = append(response.Products, p)

		response.TotalAmount += p.TotalAmount
		response.TotalDiscount += p.Discount
	}
	response.TotalAmountWithDiscount = response.TotalAmount - response.TotalDiscount
	common.Logger.LogInfo("Checkout", "checkout completed", fmt.Sprintf("items=%d", len(in.Products)))

	return response, nil
}
