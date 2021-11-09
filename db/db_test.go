package db

import (
	"context"
	"encoding/json"
	"reflect"
	"testing"

	"github.com/felipemarinho97/e-commerce/common"
	"github.com/felipemarinho97/e-commerce/config"
	"go.uber.org/atomic"
)

func TestDB_GetProduct(t *testing.T) {
	type fields struct {
		products *map[int32]*Product
	}
	type args struct {
		ctx      context.Context
		id       int32
		quantity int32
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    ProductResponse
		wantErr error
	}{
		{
			name: "successfully get a product that is on the stock",
			fields: fields{
				products: &map[int32]*Product{
					1: {
						ID:          1,
						Title:       "Product 1",
						Description: "Description 1",
						Amount:      *atomic.NewUint32(100),
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
				products: &map[int32]*Product{
					1: {
						ID:          1,
						Title:       "Product 1",
						Description: "Description 1",
						Amount:      *atomic.NewUint32(0),
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
				products: &map[int32]*Product{
					1: {
						ID:          1,
						Title:       "Product 1",
						Description: "Description 1",
						Amount:      *atomic.NewUint32(100),
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
				products: &map[int32]*Product{
					1: {
						ID:          1,
						Title:       "Product 1",
						Description: "Description 1",
						Amount:      *atomic.NewUint32(100),
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
				products: &map[int32]*Product{
					1: {
						ID:          1,
						Title:       "Product 1",
						Description: "Description 1",
						Amount:      *atomic.NewUint32(100),
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
				Products: tt.fields.products,
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

func TestNew(t *testing.T) {
	tests := []struct {
		name     string
		mockFile string
		want     Database
		wantErr  bool
	}{
		{
			name:     "successfully create a new database",
			mockFile: "./fixtures/empty.json",
			want: &DB{
				Products: &map[int32]*Product{0: {}},
				Gifts:    &map[int32]*Product{},
			},
			wantErr: false,
		},
		{
			name:     "fail to create a new database",
			mockFile: "invalid.json",
			want:     nil,
			wantErr:  true,
		},
		{
			name:     "fail to create a new database",
			mockFile: "./fixtures/fail.txt",
			want:     nil,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			config.Get().ProductsMockFile = tt.mockFile
			got, err := New()
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			g, err := json.Marshal(got)
			if err != nil {
				t.Errorf("failed to marshal got: %v", err)
			}
			w, err := json.Marshal(tt.want)
			if err != nil {
				t.Errorf("failed to marshal tt.want: %v", err)
			}
			if string(g) != string(w) {
				t.Errorf("New() = %v, want %v", string(g), string(w))
			}
		})
	}
}
