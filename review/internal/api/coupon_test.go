package api_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"coupon_service/internal/api"
	"coupon_service/internal/service/entity"
)

type MockService struct {
	mock.Mock
}

func (m *MockService) ApplyCoupon(basket entity.Basket, code string) (*entity.Basket, error) {
	args := m.Called(basket, code)
	return args.Get(0).(*entity.Basket), args.Error(1)
}

func (m *MockService) CreateCoupon(amount int, code string, expiry int) error {
	args := m.Called(amount, code, expiry)
	return args.Error(0)
}

func (m *MockService) GetCoupons(codes []string) ([]entity.Coupon, error) {
	args := m.Called(codes)
	return args.Get(0).([]entity.Coupon), args.Error(1)
}

func TestApplyCoupon_ValidRequest_ReturnsOK(t *testing.T) {
	t.Parallel()

	gin.SetMode(gin.TestMode)
	cfg := api.Config{Host: "localhost", Port: 8080}
	svc := new(MockService)
	apiInstance := api.New(cfg, svc)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	basket := entity.Basket{}
	svc.On("ApplyCoupon", basket, "TESTCODE").Return(&basket, nil)

	body := `{"basket":{},"code":"TESTCODE"}`
	c.Request, _ = http.NewRequest(http.MethodPost, "/api/apply", strings.NewReader(body))
	c.Request.Header.Set("Content-Type", "application/json")
	apiInstance.Apply(c)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestApplyCoupon_InvalidRequest_ReturnsBadRequest(t *testing.T) {
	t.Parallel()

	gin.SetMode(gin.TestMode)
	cfg := api.Config{Host: "localhost", Port: 8080}
	svc := new(MockService)
	apiInstance := api.New(cfg, svc)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request, _ = http.NewRequest(http.MethodPost, "/api/apply", nil)
	apiInstance.Apply(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestCreateCoupon_ValidRequest_ReturnsOK(t *testing.T) {
	t.Parallel()

	gin.SetMode(gin.TestMode)
	cfg := api.Config{Host: "localhost", Port: 8080}
	svc := new(MockService)
	apiInstance := api.New(cfg, svc)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	svc.On("CreateCoupon", 10, "NEWCODE", 30).Return(nil)

	body := `{"discount":10,"code":"NEWCODE","minBasketValue":30}`
	c.Request, _ = http.NewRequest(http.MethodPost, "/api/create", strings.NewReader(body))
	c.Request.Header.Set("Content-Type", "application/json")
	apiInstance.Create(c)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCreateCoupon_InvalidRequest_ReturnsBadRequest(t *testing.T) {
	t.Parallel()

	gin.SetMode(gin.TestMode)
	cfg := api.Config{Host: "localhost", Port: 8080}
	svc := new(MockService)
	apiInstance := api.New(cfg, svc)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request, _ = http.NewRequest(http.MethodPost, "/api/create", nil)
	apiInstance.Create(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetCoupons_ValidRequest_ReturnsOK(t *testing.T) {
	t.Parallel()

	gin.SetMode(gin.TestMode)
	cfg := api.Config{Host: "localhost", Port: 8080}
	svc := new(MockService)
	apiInstance := api.New(cfg, svc)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	var coupons []entity.Coupon
	svc.On("GetCoupons", []string{"CODE1", "CODE2"}).Return(coupons, nil)

	body := `{"codes":["CODE1","CODE2"]}`
	c.Request, _ = http.NewRequest(http.MethodGet, "/api/coupons", strings.NewReader(body))
	c.Request.Header.Set("Content-Type", "application/json")
	apiInstance.Get(c)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestGetCoupons_InvalidRequest_ReturnsBadRequest(t *testing.T) {
	t.Parallel()

	gin.SetMode(gin.TestMode)
	cfg := api.Config{Host: "localhost", Port: 8080}
	svc := new(MockService)
	apiInstance := api.New(cfg, svc)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request, _ = http.NewRequest(http.MethodGet, "/api/coupons", nil)
	apiInstance.Get(c)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
