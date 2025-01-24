package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"coupon_service/internal/service/entity"
)

type Service interface {
	ApplyCoupon(entity.Basket, string) (*entity.Basket, error)
	CreateCoupon(int, string, int) error
	GetCoupons([]string) ([]entity.Coupon, error)
}

type Config struct {
	Host string
	Port int
}

type API struct {
	Srv *http.Server
	Mux *gin.Engine
	svc Service
	cfg Config
}

func New[T Service](cfg Config, svc T) API {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery())

	return API{
		Mux: r,
		cfg: cfg,
		svc: svc,
	}.withServer()
}

func (a API) withServer() API {
	a.Srv = &http.Server{
		Addr:    fmt.Sprintf(":%d", a.cfg.Port),
		Handler: a.Mux,
	}

	return a
}

func (a API) withRoutes() API {
	apiGroup := a.Mux.Group("/api")
	apiGroup.POST("/apply", a.Apply)
	apiGroup.POST("/create", a.Create)
	apiGroup.GET("/coupons", a.Get)
	return a
}

func (a API) Start() error {
	a.withRoutes()

	err := a.Srv.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}

func (a API) Close() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := a.Srv.Shutdown(ctx); err != nil {
		log.Printf("Server shutdown error: %v", err)
		return
	}

	log.Println("Server shutdown")
}
