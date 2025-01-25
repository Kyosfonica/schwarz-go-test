package memdb

import (
	"fmt"
	"sync"

	"coupon_service/internal/service/entity"
)

type Repository struct {
	entries map[string]entity.Coupon
	mutex   sync.RWMutex
}

func New() *Repository {
	return &Repository{
		entries: make(map[string]entity.Coupon),
	}
}

func (r *Repository) FindByCode(code string) (*entity.Coupon, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	coupon, ok := r.entries[code]
	if !ok {
		return nil, fmt.Errorf("coupon not found")
	}

	return &coupon, nil
}

func (r *Repository) Save(coupon entity.Coupon) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.entries[coupon.Code] = coupon

	return nil
}
