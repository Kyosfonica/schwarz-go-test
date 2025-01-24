package service

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/mock"

	"coupon_service/internal/service/entity"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) FindByCode(code string) (*entity.Coupon, error) {
	args := m.Called(code)
	return args.Get(0).(*entity.Coupon), args.Error(1)
}

func (m *MockRepository) Save(coupon entity.Coupon) error {
	args := m.Called(coupon)
	return args.Error(0)
}

func TestNew(t *testing.T) {
	t.Parallel()

	type args struct {
		repo Repository
	}
	tests := []struct {
		name string
		args args
		want Service
	}{
		{"initialize service", args{repo: nil}, Service{repo: nil}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.repo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestService_ApplyCoupon(t *testing.T) {
	t.Parallel()

	mockRepo := new(MockRepository)
	mockRepo.On("FindByCode", "VALIDCOUPON").Return(&entity.Coupon{Discount: 10}, nil)
	mockRepo.On("FindByCode", "INVALIDCOUPON").Return(&entity.Coupon{}, fmt.Errorf("coupon not found"))

	type fields struct {
		repo Repository
	}
	type args struct {
		basket entity.Basket
		code   string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantB   *entity.Basket
		wantErr bool
	}{
		{
			name:   "apply valid coupon",
			fields: fields{repo: mockRepo},
			args: args{
				basket: entity.Basket{Value: 100},
				code:   "VALIDCOUPON",
			},
			wantB:   &entity.Basket{Value: 100, AppliedDiscount: 10, ApplicationSuccessful: true},
			wantErr: false,
		},
		{
			name:   "apply invalid coupon",
			fields: fields{repo: mockRepo},
			args: args{
				basket: entity.Basket{Value: 100},
				code:   "INVALIDCOUPON",
			},
			wantB:   nil,
			wantErr: true,
		},
		{
			name:   "apply coupon to zero value basket",
			fields: fields{repo: mockRepo},
			args: args{
				basket: entity.Basket{Value: 0},
				code:   "VALIDCOUPON",
			},
			wantB:   nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Service{
				repo: tt.fields.repo,
			}
			gotB, err := s.ApplyCoupon(tt.args.basket, tt.args.code)
			if (err != nil) != tt.wantErr {
				t.Errorf("ApplyCoupon() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotB, tt.wantB) {
				t.Errorf("ApplyCoupon() gotB = %v, want %v", gotB, tt.wantB)
			}
		})
	}
}

func TestService_CreateCoupon(t *testing.T) {
	t.Parallel()

	mockRepo := new(MockRepository)
	mockRepo.On("Save", mock.Anything).Return(nil)

	type fields struct {
		repo Repository
	}
	type args struct {
		discount       int
		code           string
		minBasketValue int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"Apply 10%", fields{mockRepo}, args{10, "Superdiscount", 55}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Service{
				repo: tt.fields.repo,
			}

			err := s.CreateCoupon(tt.args.discount, tt.args.code, tt.args.minBasketValue)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateCoupon() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
