package discount

import (
	"context"

	"github.com/felipemarinho97/e-commerce/config"
	"google.golang.org/grpc"
)

// ClientConn useful for testing.
type ClientConn interface {
	Close() error
}

// DiscountService is the gRPC client for the Discount service.
type DiscountService struct {
	conn   ClientConn
	client DiscountClient
}

// New returns a new DiscountService.
var New = func(addr string) (*DiscountService, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return &DiscountService{}, err
	}

	client := NewDiscountClient(conn)

	return &DiscountService{
		conn:   conn,
		client: client,
	}, nil
}

// Quit closes the connection.
func (dc *DiscountService) Quit() (err error) {
	err = dc.conn.Close()
	if err != nil {
		return
	}

	return
}

// GetDiscountPercentage returns the discount percentage for a given product.
func GetDiscountPercentage(ctx context.Context, productID int32) (float32, error) {
	dc, err := New(config.Get().DiscountAddr)
	if err != nil {
		return 0, err
	}
	defer dc.Quit()

	out, err := dc.client.GetDiscount(ctx, &GetDiscountRequest{
		ProductID: productID,
	})
	if err != nil {
		return 0, err
	}

	return out.Percentage, nil
}
