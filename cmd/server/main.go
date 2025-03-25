package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/swaggo/echo-swagger"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"song-lib/internal/config"
	"song-lib/internal/db"
	"song-lib/internal/externalAPI"
	"song-lib/internal/handlers"
	"song-lib/internal/repository/postgres"
	"song-lib/internal/usecase"
	"syscall"
	"time"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("failed to initialize logger: %v", err)
	}
	defer func(logger *zap.Logger) {
		err := logger.Sync()
		if err != nil {

		}
	}(logger)

	sugar := logger.Sugar()

	err = config.SetUp()
	if err != nil {
		sugar.Fatalw("failed to fetch config", "error", err)
	}
	postgresDB, err := db.InitDB()
	if err != nil {
		sugar.Fatalw("failed to initialize database", "error", err)
	}
	err = db.MakeMigrations(true)
	if err != nil {
		sugar.Fatalw("failed to make migrations", "error", err)
	}

	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	myClient := externalAPI.NewClient(
		fmt.Sprintf("http://%s:%s",
			config.AppConfig.Server.Host,
			config.AppConfig.Server.Port)) // внешнее апи /info находится по пути http://localhost:8080/info, изменить, если требуется

	songRepo := postgres.NewSongRepo(postgresDB, sugar)
	songUseCase := usecase.NewSongInstance(songRepo, *myClient, sugar)
	songHandlers := handlers.NewSongHandler(songUseCase, sugar)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	songGroup := e.Group("/api/songs")

	songGroup.POST("", songHandlers.Create)
	songGroup.GET("/:id", songHandlers.Get)
	songGroup.GET("/filter", songHandlers.GetSongs)
	songGroup.PUT("/:id", songHandlers.Update)
	songGroup.DELETE("/:id", songHandlers.Delete)

	sugar.Infow("starting server", "port", 8080)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := e.Start(":8080"); err != nil && !errors.Is(err, http.ErrServerClosed) {
			sugar.Fatalw("failed to start server", "error", err)
		}
	}()

	<-stop
	sugar.Infow("received shutdown signal, starting shutdown...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		sugar.Fatalw("failed to gracefully shut down server", "error", err)
	}

	sugar.Infow("server gracefully stopped")
}
