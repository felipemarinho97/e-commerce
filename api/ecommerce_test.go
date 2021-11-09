package api

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/felipemarinho97/e-commerce/common"
	"github.com/felipemarinho97/e-commerce/db"
	pb "github.com/felipemarinho97/e-commerce/examples/go/protos/api"
	"go.uber.org/atomic"
)

func Test_server_Checkout(t *testing.T) {
	type fields struct {
		UnimplementedEcommerceServiceServer pb.UnimplementedEcommerceServiceServer
		db                                  db.Database
	}
	type args struct {
		ctx context.Context
		in  *pb.CheckoutRequest
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *pb.CheckoutResponse
		wantErr error
	}{
		{
			name: "succesful checkout: product in stock",
			fields: fields{
				db: &db.DB{
					Products: &map[int64]*db.Product{
						1: {
							ID:          1,
							Title:       "Product 1",
							Description: "Product 1 description",
							Amount:      *atomic.NewUint64(100),
							Price:       15654,
							IsGift:      false,
						},
					},
				},
			},
			args: args{
				ctx: context.Background(),
				in: &pb.CheckoutRequest{
					Products: []*pb.ProductRequest{
						{
							Id:       1,
							Quantity: 1,
						},
					},
				},
			},
			want: &pb.CheckoutResponse{
				Products: []*pb.ProductResponse{
					{
						Id:          1,
						Quantity:    1,
						UnitAmount:  15654,
						TotalAmount: 15654,
						Discount:    0,
						IsGift:      false,
					},
				},
				TotalAmount:             15654,
				TotalAmountWithDiscount: 15654,
				TotalDiscount:           0,
			},
			wantErr: nil,
		},
		{
			name: "succesful checkout: two products",
			fields: fields{
				db: &db.DB{
					Products: &map[int64]*db.Product{
						1: {
							ID:          1,
							Title:       "Product 1",
							Description: "Product 1 description",
							Amount:      *atomic.NewUint64(100),
							Price:       15654,
							IsGift:      false,
						},
						2: {
							ID:          2,
							Title:       "Product 2",
							Description: "Product 2 description",
							Amount:      *atomic.NewUint64(100),
							Price:       15654,
							IsGift:      false,
						},
					},
				},
			},
			args: args{
				ctx: context.Background(),
				in: &pb.CheckoutRequest{
					Products: []*pb.ProductRequest{
						{
							Id:       1,
							Quantity: 1,
						},
						{
							Id:       2,
							Quantity: 1,
						},
					},
				},
			},
			want: &pb.CheckoutResponse{
				Products: []*pb.ProductResponse{
					{
						Id:          1,
						Quantity:    1,
						UnitAmount:  15654,
						TotalAmount: 15654,
						Discount:    0,
						IsGift:      false,
					},
					{
						Id:          2,
						Quantity:    1,
						UnitAmount:  15654,
						TotalAmount: 15654,
						Discount:    0,
						IsGift:      false,
					},
				},
				TotalAmount:             31308,
				TotalAmountWithDiscount: 31308,
				TotalDiscount:           0,
			},
			wantErr: nil,
		},
		{
			name: "error on checkout: product not in stock",
			fields: fields{
				db: &db.DB{
					Products: &map[int64]*db.Product{
						1: {
							ID:          1,
							Title:       "Product 1",
							Description: "Product 1 description",
							Amount:      *atomic.NewUint64(0),
							Price:       15654,
							IsGift:      false,
						},
					},
				},
			},
			args: args{
				ctx: context.Background(),
				in: &pb.CheckoutRequest{
					Products: []*pb.ProductRequest{
						{
							Id:       1,
							Quantity: 1,
						},
					},
				},
			},
			want: &pb.CheckoutResponse{
				Products: []*pb.ProductResponse{
					{
						Id:          1,
						Quantity:    -1,
						UnitAmount:  0,
						TotalAmount: 0,
						Discount:    0,
						IsGift:      false,
					},
				},
			},
			wantErr: nil,
		},
		{
			name: "error on checkout: product not in database",
			fields: fields{
				db: &db.DB{
					Products: &map[int64]*db.Product{
						1: {
							ID:          1,
							Title:       "Product 1",
							Description: "Product 1 description",
							Amount:      *atomic.NewUint64(100),
							Price:       15654,
							IsGift:      false,
						},
					},
				},
			},
			args: args{
				ctx: context.Background(),
				in: &pb.CheckoutRequest{
					Products: []*pb.ProductRequest{
						{
							Id:       2,
							Quantity: 1,
						},
					},
				},
			},
			want:    nil,
			wantErr: fmt.Errorf("error checking out product with id=2: %s", common.ErrProductNotFound.Error()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := server{
				UnimplementedEcommerceServiceServer: tt.fields.UnimplementedEcommerceServiceServer,
				db:                                  tt.fields.db,
			}
			got, err := s.Checkout(tt.args.ctx, tt.args.in)
			if err != nil && err.Error() != tt.wantErr.Error() {
				t.Errorf("server.Checkout() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("server.Checkout() = %v,\n want %v", got, tt.want)
			}
		})
	}
}
