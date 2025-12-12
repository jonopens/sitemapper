package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"jonopens/sitemapper/internal/config"
	"jonopens/sitemapper/internal/database"
	"jonopens/sitemapper/internal/handlers"
	"jonopens/sitemapper/internal/services"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig("configs/config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize database
	db, err := database.NewDatabase(cfg)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	// Initialize services
	sitemapService := services.NewSitemapService(db)
	reportService := services.NewReportService(db)
	_ = services.NewJobService(db)         // TODO: Use when job endpoints are implemented
	_ = services.NewGroupingService(db)    // TODO: Use when grouping endpoints are implemented

	// Initialize handlers
	sitemapHandler := handlers.NewSitemapHandler(sitemapService)
	reportHandler := handlers.NewReportHandler(reportService)
	userHandler := handlers.NewUserHandler(db)

	// Setup router
	router := gin.Default()

	// Setup routes
	setupRoutes(router, sitemapHandler, reportHandler, userHandler)

	// Create server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.ServerPort),
		Handler: router,
	}

	// Start server in goroutine
	go func() {
		log.Printf("Starting server on port %d", cfg.ServerPort)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exited")
}

func setupRoutes(router *gin.Engine, sitemapHandler *handlers.SitemapHandler, reportHandler *handlers.ReportHandler, userHandler *handlers.UserHandler) {
	api := router.Group("/api/v1")
	{
		// Sitemap routes
		api.POST("/sitemaps", sitemapHandler.Upload)
		api.GET("/sitemaps/:id", sitemapHandler.Get)
		api.GET("/sitemaps", sitemapHandler.List)

		// Report routes
		api.GET("/reports/:id", reportHandler.Get)
		api.GET("/reports", reportHandler.List)
		api.POST("/reports/:id/generate", reportHandler.Generate)

		// User routes
		api.GET("/users/:id", userHandler.Get)
		api.POST("/users", userHandler.Create)
	}
}

