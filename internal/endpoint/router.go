package endpoint

import (
	"net/http"

	"github.com/Gierdiaz/diagier-clinics/internal/setup"
	"github.com/Gierdiaz/diagier-clinics/pkg/messaging"
	_ "github.com/Gierdiaz/diagier-clinics/pkg/middleware"
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

		v1.GET("/patients", patientHandler.GetAllPatients)
		v1.GET("/patients/:id", patientHandler.GetPatientByID)
		v1.POST("/patients", patientHandler.CreatePatient)
		v1.PUT("/patients/:id", patientHandler.UpdatePatient)
		v1.DELETE("/patients/:id", patientHandler.DeletePatient)

		// Rotas protegidas
		// authorized := v1.Group("/")
		// authorized.Use(middleware.AuthMiddleware())
		// {
		// 	// Rotas para pacientes
		// 	authorized.GET("patients", patientHandler.GetAllPatients)
		// 	authorized.GET("patients/:id", patientHandler.GetPatientByID)
		// 	authorized.POST("patients", patientHandler.CreatePatient)
		// 	authorized.PUT("patients/:id", patientHandler.UpdatePatient)
		// 	authorized.DELETE("patients/:id", patientHandler.DeletePatient)
		// }
	}

	return router
}
