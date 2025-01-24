package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"coupon_service/internal/api/entity"
)

func (a API) Apply(c *gin.Context) {
	apiReq := entity.ApplicationRequest{}
	if err := c.ShouldBindJSON(&apiReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	basket, err := a.svc.ApplyCoupon(apiReq.Basket, apiReq.Code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, basket)
}

func (a API) Create(c *gin.Context) {
	apiReq := entity.Coupon{}
	if err := c.ShouldBindJSON(&apiReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := a.svc.CreateCoupon(apiReq.Discount, apiReq.Code, apiReq.MinBasketValue)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusOK)
}

func (a API) Get(c *gin.Context) {
	apiReq := entity.CouponRequest{}
	if err := c.ShouldBindJSON(&apiReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	coupons, err := a.svc.GetCoupons(apiReq.Codes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, coupons)
}
