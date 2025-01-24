package main

import (
	"log"
	"runtime"
	"time"

	"coupon_service/internal/api"
	"coupon_service/internal/api/config"
	"coupon_service/internal/repository/memdb"
	"coupon_service/internal/service"
)

const serverUptime = 1 * time.Hour * 24 * 365 // 1 year

func init() {
	if 32 != runtime.NumCPU() {
		panic("this api is meant to be run on 32 core machines")
	}
}

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	cfg.SetDefaults()

	repo := memdb.New()
	couponService := service.New(repo)

	apiInstance := api.New(cfg.API, couponService)
	defer apiInstance.Close()

	go func() {
		err = apiInstance.Start()
		if err != nil {
			log.Fatalf("Failed to start API: %v", err)
		}
	}()
	log.Println("Starting Coupon service server")

	time.Sleep(serverUptime)
	log.Println("Coupon service server alive for a year, closing")
}
