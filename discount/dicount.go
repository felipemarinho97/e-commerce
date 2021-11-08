package discount

import (
	"context"
	"log"

	"github.com/felipemarinho97/e-commerce/config"
	"google.golang.org/grpc"
)

type DiscountService struct {
	conn   *grpc.ClientConn
	client *DiscountClient
}

func New(addr string) (*DiscountService, error) {
	conn, err := grpc.Dial(addr)
	if err != nil {
		return &DiscountService{}, err
	}

	client := NewDiscountClient(conn)

	return &DiscountService{
		conn:   conn,
		client: &client,
	}, nil
}

func (dc *DiscountService) Quit() (err error) {
	err = dc.conn.Close()
	if err != nil {
		return
	}

	return
}

func GetDiscountPercentage(ctx context.Context, productID int32) (float32, error) {
	dc, err := New(config.Get().DiscountAddr)
	if err != nil {
		log.Fatalf("Error")
	}
	defer dc.Quit()

	out, err := (*dc.client).GetDiscount(ctx, &GetDiscountRequest{
		ProductID: productID,
	})
	if err != nil {
		return 0, err
	}

	return out.Percentage, nil
}
