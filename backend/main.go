package main

import (
	"daily-worker-roster-management-system/db"
	handler "daily-worker-roster-management-system/handlers"
	"daily-worker-roster-management-system/repository"
	"daily-worker-roster-management-system/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	database := db.InitDB("roster_data/roster.db") // Use volume path
	db.Migrate(database)

	// Setup repos, usecases, handlers
	shiftRepo := repository.NewShiftRepository(database)
	shiftRequestRepo := repository.NewShiftRequestRepository(database)
	assignmentRepo := repository.NewAssignmentRepository(database)

	shiftUsecase := usecase.NewShiftUsecase(shiftRepo)
	userUsecase := usecase.NewUserUsecase()
	shiftRequestUsecase := usecase.NewShiftRequestUsecase(shiftRequestRepo, shiftRepo, assignmentRepo)
	assignmentUsecase := usecase.NewAssignmentUsecase(assignmentRepo)

	shiftHandler := handler.NewShiftHandler(shiftUsecase)
	shiftRequestHandler := handler.NewShiftRequestHandler(shiftRequestUsecase)
	assignmentHandler := handler.NewAssignmentHandler(assignmentUsecase)
	authHandler := handler.NewAuthHandler(userUsecase, []byte("your-very-secret-key"))

	r := gin.Default()
	r.Use(CORSMiddleware())
	r.GET("/shifts", shiftHandler.GetShifts)
	r.POST("/shifts", shiftHandler.CreateShift)

	r.POST("/requests", shiftRequestHandler.Create)
	r.GET("/requests/pending", shiftRequestHandler.ListPending)
	r.PUT("/requests/:id/approve", shiftRequestHandler.Approve)
	r.PUT("/requests/:id/reject", shiftRequestHandler.Reject)

	r.POST("/assignments", assignmentHandler.Create)
	r.GET("/assignments", assignmentHandler.Get)

	r.POST("/auth/login", authHandler.Login)
	r.GET("/auth/verify", authHandler.VerifyToken)

	r.Run(":9400")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
