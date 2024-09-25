package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/zinrai/alert-hub-go/internal/db"
	"github.com/zinrai/alert-hub-go/internal/handler"
	"github.com/zinrai/alert-hub-go/internal/repository"
	"github.com/zinrai/alert-hub-go/internal/usecase"
)

func main() {
	dbConn, err := db.NewConnection()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer dbConn.Close()

	alertRepo, err := repository.NewPostgresAlertRepository(dbConn)
	if err != nil {
		log.Fatalf("Failed to create alert repository: %v", err)
	}

	alertUsecase := usecase.NewAlertUsecase(alertRepo)
	alertHandler := handler.NewAlertHandler(alertUsecase)

	r := gin.Default()

	api := r.Group("/api")
	{
		api.GET("/alerts", alertHandler.GetAlerts)
		api.POST("/alerts", alertHandler.CreateAlert)
		api.GET("/alerts/:id", alertHandler.GetAlert)
		api.PATCH("/alerts/:id", alertHandler.UpdateAlert)
	}

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
