package main

import (
	"log"

	"ovo-gaming/server/internal/config"
	"ovo-gaming/server/internal/db"
	"ovo-gaming/server/internal/handler"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	cfg := config.Load()
	var gdb = mustOpenDB(cfg)

	r := gin.Default()
	api := r.Group("/api")
	{
		api.GET("/health", handler.Health)

		mini := api.Group("/mini")
		miniHandler := handler.New(gdb)
		mini.GET("/home", miniHandler.MiniHome)

		admin := api.Group("/admin")
		adminHandler := handler.New(gdb)
		admin.GET("/banners", adminHandler.AdminListBanners)
		admin.POST("/banners", adminHandler.AdminCreateBanner)
		admin.PUT("/banners/:id", adminHandler.AdminUpdateBanner)
		admin.GET("/notices", adminHandler.AdminListNotices)
		admin.POST("/notices", adminHandler.AdminCreateNotice)
		admin.GET("/packages", adminHandler.AdminListPackages)
		admin.POST("/packages", adminHandler.AdminCreatePackage)
		admin.PUT("/packages/:id", adminHandler.AdminUpdatePackage)
	}

	log.Printf("server listening on %s", cfg.HTTPAddr)
	if err := r.Run(cfg.HTTPAddr); err != nil {
		log.Fatal(err)
	}
}

func mustOpenDB(cfg config.Config) *gorm.DB {
	gdb, err := db.Open(cfg)
	if err != nil {
		log.Fatal(err)
	}
	return gdb
}
