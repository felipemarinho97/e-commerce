package discount

import (
	"context"
	"errors"
	"testing"

	"google.golang.org/grpc"
)

// MockClientConn is a mock implementation of grpc.ClientConn.
type MockClientConn struct{}

// Close the mock client connection.
func (cc *MockClientConn) Close() error {
	return nil
}

// MockDiscountClient is a mock implementation of DiscountClient.
type MockDiscountClient struct {
	GetDiscountMock func(ctx context.Context, in *GetDiscountRequest, opts ...grpc.CallOption) (*GetDiscountResponse, error)
}

// GetDiscount is a mock implementation of DiscountClient.GetDiscount.
func (mdc *MockDiscountClient) GetDiscount(ctx context.Context, in *GetDiscountRequest, opts ...grpc.CallOption) (*GetDiscountResponse, error) {
	return mdc.GetDiscountMock(ctx, in, opts...)
}

// MockGetPercentage mocks the implementation of GetPercentage.
func MockGetPercentage(p float32, wantErr bool) {
	mockClient := &MockDiscountClient{
		GetDiscountMock: func(ctx context.Context, in *GetDiscountRequest, opts ...grpc.CallOption) (*GetDiscountResponse, error) {
			if wantErr {
				return nil, errors.New("network error")
			}
			return &GetDiscountResponse{
				Percentage: p,
			}, nil
		},
	}
	// mock DiscountService
	New = func(addr string) (*DiscountService, error) {
		return &DiscountService{
			conn:   &MockClientConn{},
			client: mockClient,
		}, nil
	}
}

func TestGetDiscountPercentage(t *testing.T) {
	type args struct {
		ctx       context.Context
		productID int32
	}
	tests := []struct {
		name    string
		args    args
		want    float32
		wantErr bool
	}{
		{
			name: "test success return",
			args: args{
				ctx:       context.Background(),
				productID: 1,
			},
			want:    0.1,
			wantErr: false,
		},
		{
			name: "test error return",
			args: args{
				ctx:       context.Background(),
				productID: 1,
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MockGetPercentage(tt.want, tt.wantErr)
			got, err := GetDiscountPercentage(tt.args.ctx, tt.args.productID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDiscountPercentage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetDiscountPercentage() = %v, want %v", got, tt.want)
			}
		})
	}
}
