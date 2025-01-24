package service

import (
	"fmt"

	"github.com/google/uuid"

	"coupon_service/internal/service/entity"
)

type Repository interface {
	FindByCode(string) (*entity.Coupon, error)
	Save(entity.Coupon) error
}

type Service struct {
	repo Repository
}

func New(repo Repository) Service {
	return Service{
		repo: repo,
	}
}

func (s Service) ApplyCoupon(basket entity.Basket, code string) (b *entity.Basket, e error) {
	b = &basket
	coupon, err := s.repo.FindByCode(code)
	if err != nil {
		return nil, err
	}

	if b.Value > 0 {
		b.AppliedDiscount = coupon.Discount
		b.ApplicationSuccessful = true

		return b, nil
	}
	if b.Value == 0 {
		return nil, fmt.Errorf("basket value is zero")
	}

	return nil, fmt.Errorf("tried to apply discount to negative value")
}

func (s Service) CreateCoupon(discount int, code string, minBasketValue int) error {
	coupon := entity.Coupon{
		Discount:       discount,
		Code:           code,
		MinBasketValue: minBasketValue,
		ID:             uuid.NewString(),
	}

	err := s.repo.Save(coupon)
	if err != nil {
		return err
	}

	return nil
}

func (s Service) GetCoupons(codes []string) ([]entity.Coupon, error) {
	coupons := make([]entity.Coupon, 0, len(codes))
	var e error = nil

	for idx, code := range codes {
		coupon, err := s.repo.FindByCode(code)
		if err != nil {
			if e == nil {
				e = fmt.Errorf("code: %s, index: %d", code, idx)
			} else {
				e = fmt.Errorf("%w; code: %s, index: %d", e, code, idx)
			}

			continue
		}
		coupons = append(coupons, *coupon)
	}

	return coupons, e
}
