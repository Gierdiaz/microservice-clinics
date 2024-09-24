package endpoint

import (
	"net/http"

	"github.com/Gierdiaz/diagier-clinics/internal/setup"
	"github.com/Gierdiaz/diagier-clinics/pkg/messaging"
	"github.com/Gierdiaz/diagier-clinics/pkg/middleware"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func Router(db *sqlx.DB, rabbitMQ *messaging.RabbitMQ) *gin.Engine {
	router := gin.Default()

	patientHandler := setup.SetupServices(db, rabbitMQ)
	userHandler := setup.SetupUserServices(db)

	v1 := router.Group("/api/v1")
	{
		// Rota de health check
		v1.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "OK",
			})
		})

		// Rotas de autenticação
		v1.POST("/register", userHandler.Register)
		v1.POST("/login", userHandler.Login)

		v1.GET("/patients", middleware.AuthMiddleware(), patientHandler.GetAllPatients)
		v1.GET("/patients/:id", middleware.AuthMiddleware(), patientHandler.GetPatientByID)
		v1.POST("/patients", middleware.AuthMiddleware(), patientHandler.CreatePatient)
		v1.PUT("/patients/:id", middleware.AuthMiddleware(), patientHandler.UpdatePatient)
		v1.DELETE("/patients/:id", middleware.AuthMiddleware(), patientHandler.DeletePatient)

	}

	return router
}
