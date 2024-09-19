package endpoint

import (
	"net/http"

	"github.com/Gierdiaz/diagier-clinics/internal/setup"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func Router(db *sqlx.DB) *gin.Engine {
	router := gin.Default()

	patientHandler := setup.SetupServices(db)

	v1 := router.Group("/api/v1")
	{
		// Rota de health check
		v1.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "OK",
			})
		})


		// Rotas para pacientes
		v1.GET("/patients", patientHandler.GetAllPatients)
		v1.GET("/patients/:id", patientHandler.GetPatientByID)
		v1.POST("/patients", patientHandler.CreatePatient)
		v1.PUT("/patients/:id", patientHandler.UpdatePatient)
		v1.DELETE("/patients/:id", patientHandler.DeletePatient)
	}

	return router
}
