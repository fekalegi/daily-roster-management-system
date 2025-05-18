package main

import (
	"daily-worker-roster-management-system/db"
	handler "daily-worker-roster-management-system/handlers"
	"daily-worker-roster-management-system/middleware"
	"daily-worker-roster-management-system/repository"
	"daily-worker-roster-management-system/usecase"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	database := db.InitDB("roster_data/roster.db") // Use volume path
	db.Migrate(database)

	// Setup repos, usecases, handlers
	shiftRepo := repository.NewShiftRepository(database)
	shiftRequestRepo := repository.NewShiftRequestRepository(database)
	assignmentRepo := repository.NewAssignmentRepository(database)
	workerRepo := repository.NewWorkerRequestRepository(database)

	shiftUsecase := usecase.NewShiftUsecase(shiftRepo)
	userUsecase := usecase.NewUserUsecase(workerRepo)
	shiftRequestUsecase := usecase.NewShiftRequestUsecase(shiftRequestRepo, shiftRepo, assignmentRepo)
	assignmentUsecase := usecase.NewAssignmentUsecase(assignmentRepo, shiftRepo)
	workerUsecase := usecase.NewWorkerUsecase(workerRepo)

	shiftHandler := handler.NewShiftHandler(shiftUsecase)
	shiftRequestHandler := handler.NewShiftRequestHandler(shiftRequestUsecase)
	assignmentHandler := handler.NewAssignmentHandler(assignmentUsecase)
	workerHandler := handler.NewWorkerHandler(workerUsecase)
	secretJwt := os.Getenv("JWT_SECRET")
	authHandler := handler.NewAuthHandler(userUsecase, []byte(secretJwt))

	r := gin.Default()
	r.Use(CORSMiddleware())

	r.POST("/auth/login", authHandler.Login)
	r.GET("/auth/verify", authHandler.VerifyToken)

	r.Use(middleware.JWTAuthMiddleware([]byte(secretJwt)))
	// Admin-only routes

	admin := r.Group("/")
	admin.Use(middleware.AdminOnlyMiddleware())
	admin.POST("/shifts", shiftHandler.CreateShift)
	r.GET("/shifts", shiftHandler.GetShifts)

	r.POST("/requests", shiftRequestHandler.Create)
	r.GET("/requests/pending", shiftRequestHandler.ListPending)
	admin.PUT("/requests/:id/approve", shiftRequestHandler.Approve)
	admin.PUT("/requests/:id/reject", shiftRequestHandler.Reject)

	r.GET("/assignments", assignmentHandler.Get)
	admin.POST("/assignments", assignmentHandler.Create)

	admin.GET("/workers", workerHandler.GetAll)

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
