package main

import (
	"log/slog"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kenueyy/smtp-service/internal/handlers"
)

var logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

func main() {
	r := gin.Default()
	v1 := r.Group("api/v1")
	v1.POST("/sendcode", handlers.SendCode)

	logger.Info("server starting",
		"port", 4444,
	)

	if err := r.Run(":" + strconv.Itoa(4444)); err != nil {
		logger.Error("server stopped with error",
			"error", err.Error(),
		)
		os.Exit(1)
	}
}
