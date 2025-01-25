package memdb_test

import (
	"coupon_service/internal/repository/memdb"
	"coupon_service/internal/service/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepository_FindByCode_Success(t *testing.T) {
	t.Parallel()

	repo := memdb.New()
	coupon := entity.Coupon{Code: "SAVE10"}
	err := repo.Save(coupon)
	assert.NoError(t, err)

	result, err := repo.FindByCode("SAVE10")
	assert.NoError(t, err)
	assert.Equal(t, "SAVE10", result.Code)
}

func TestRepository_FindByCode_NotFound(t *testing.T) {
	t.Parallel()

	repo := memdb.New()

	_, err := repo.FindByCode("NOTEXIST")
	assert.Error(t, err)
	assert.Equal(t, "coupon not found", err.Error())
}

func TestRepository_Save_Success(t *testing.T) {
	t.Parallel()

	repo := memdb.New()
	coupon := entity.Coupon{Code: "SAVE20"}

	err := repo.Save(coupon)
	assert.NoError(t, err)

	result, err := repo.FindByCode("SAVE20")
	assert.NoError(t, err)
	assert.Equal(t, "SAVE20", result.Code)
}
