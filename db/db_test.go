package db

import (
	"context"
	"reflect"
	"testing"

	"github.com/felipemarinho97/e-commerce/common"
	"go.uber.org/atomic"
)

func TestDB_GetProduct(t *testing.T) {
	type fields struct {
		products *map[int64]*Product
	}
	type args struct {
		ctx      context.Context
		id       int64
		quantity int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    ProductResponse
		wantErr error
	}{
		{
			name: "successfuly get a product that is on the stock",
			fields: fields{
				products: &map[int64]*Product{
					1: {
						ID:          1,
						Title:       "Product 1",
						Description: "Description 1",
						Amount:      *atomic.NewUint64(100),
						Price:       10,
						IsGift:      false,
					},
				},
			},
			args: args{
				ctx:      context.Background(),
				id:       1,
				quantity: 1,
			},
			want: ProductResponse{
				ID:     1,
				Amount: 1,
				Price:  10,
			},
			wantErr: nil,
		},
		{
			name: "fail to get a product that is not on the stock",
			fields: fields{
				products: &map[int64]*Product{
					1: {
						ID:          1,
						Title:       "Product 1",
						Description: "Description 1",
						Amount:      *atomic.NewUint64(0),
						Price:       10,
						IsGift:      false,
					},
				},
			},
			args: args{
				ctx:      context.Background(),
				id:       1,
				quantity: 1,
			},
			want:    ProductResponse{},
			wantErr: common.ErrProductOutOfStock,
		},
		{
			name: "fail to get a product that do not exists",
			fields: fields{
				products: &map[int64]*Product{
					1: {
						ID:          1,
						Title:       "Product 1",
						Description: "Description 1",
						Amount:      *atomic.NewUint64(100),
						Price:       10,
						IsGift:      false,
					},
				},
			},
			args: args{
				ctx:      context.Background(),
				id:       2,
				quantity: 1,
			},
			want:    ProductResponse{},
			wantErr: common.ErrProductNotFound,
		},
		{
			name: "fail to get a product that is gift",
			fields: fields{
				products: &map[int64]*Product{
					1: {
						ID:          1,
						Title:       "Product 1",
						Description: "Description 1",
						Amount:      *atomic.NewUint64(100),
						Price:       10,
						IsGift:      true,
					},
				},
			},
			args: args{
				ctx:      context.Background(),
				id:       1,
				quantity: 1,
			},
			want:    ProductResponse{},
			wantErr: common.ErrProductIsGift,
		},
		{
			name: "partially get a product that is on the stock",
			fields: fields{
				products: &map[int64]*Product{
					1: {
						ID:          1,
						Title:       "Product 1",
						Description: "Description 1",
						Amount:      *atomic.NewUint64(100),
						Price:       10,
						IsGift:      false,
					},
				},
			},
			args: args{
				ctx:      context.Background(),
				id:       1,
				quantity: 102,
			},
			want: ProductResponse{
				ID:     1,
				Amount: 100,
				Price:  10,
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := &DB{
				products: tt.fields.products,
			}
			got, err := db.GetProduct(tt.args.ctx, tt.args.id, tt.args.quantity)
			if err != tt.wantErr {
				t.Errorf("DB.GetProduct() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DB.GetProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
