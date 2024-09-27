package handler

import (
	"net/http"

	"github.com/Gierdiaz/diagier-clinics/internal/domain/patient"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type PatientsHandler struct {
	service patient.PatientService
}

func NewPatientsHandler(service patient.PatientService) *PatientsHandler {
	return &PatientsHandler{service: service}
}

func (handler *PatientsHandler) GetAllPatients(c *gin.Context) {
	patients, err := handler.service.GetAllPatients()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, patients)
}

func (handler *PatientsHandler) GetPatientByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	patient, err := handler.service.GetPatientByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, patient)
}

func (handler *PatientsHandler) CreatePatient(c *gin.Context) {
	var patient patient.Patient
	if err := c.ShouldBindJSON(&patient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdPatient, err := handler.service.CreatePatient(&patient)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, createdPatient)
}

func (handler *PatientsHandler) UpdatePatient(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var patient patient.Patient
	if err := c.ShouldBindJSON(&patient); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	patient.ID = id

	updatedPatient, err := handler.service.UpdatePatient(id, &patient)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedPatient)
}

func (handler *PatientsHandler) DeletePatient(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := handler.service.DeletePatient(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
