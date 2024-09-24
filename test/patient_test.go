package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Gierdiaz/diagier-clinics/internal/domain/patient"
	"github.com/Gierdiaz/diagier-clinics/internal/handler"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

// mockPatientService simula as operações do serviço de pacientes
type mockPatientService struct{}

func (m *mockPatientService) GetAllPatients() ([]*patient.Patient, error) {
	return []*patient.Patient{
		{
			ID:   uuid.New(),
			Name: "John Doe",
			Age:  35,
		},
	}, nil
}

func (m *mockPatientService) GetPatientByID(id uuid.UUID) (*patient.Patient, error) {
	return &patient.Patient{
		ID:   id,
		Name: "John Doe",
		Age:  35,
	}, nil
}

func (m *mockPatientService) CreatePatient(dto *patient.PatientDTO) (*patient.Patient, error) {
	return &patient.Patient{
		ID:   uuid.New(),
		Name: dto.Name,
		Age:  dto.Age,
	}, nil
}

func (m *mockPatientService) UpdatePatient(id uuid.UUID, dto *patient.PatientDTO) (*patient.Patient, error) {
	return &patient.Patient{
		ID:   id,
		Name: dto.Name,
		Age:  dto.Age,
	}, nil
}

func (m *mockPatientService) DeletePatient(id uuid.UUID) error {
	return nil
}

func TestGetAllPatients(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	patientService := &mockPatientService{}
	handler := handler.NewPatientsHandler(patientService)

	r.GET("/patients", handler.GetAllPatients)

	req, _ := http.NewRequest(http.MethodGet, "/patients", nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "John Doe")
}

func TestGetPatientByID(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	patientService := &mockPatientService{}
	handler := handler.NewPatientsHandler(patientService)

	id := uuid.New()
	r.GET("/patients/:id", handler.GetPatientByID)

	req, _ := http.NewRequest(http.MethodGet, "/patients/"+id.String(), nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "John Doe")
}

func TestCreatePatient(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	patientService := &mockPatientService{}
	handler := handler.NewPatientsHandler(patientService)

	r.POST("/patients", handler.CreatePatient)

	patientDTO := &patient.PatientDTO{
		Name: "Jane Doe",
		Age:  28,
	}
	body, _ := json.Marshal(patientDTO)

	req, _ := http.NewRequest(http.MethodPost, "/patients", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "Jane Doe")
}

func TestUpdatePatient(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	patientService := &mockPatientService{}
	handler := handler.NewPatientsHandler(patientService)

	id := uuid.New()
	r.PUT("/patients/:id", handler.UpdatePatient)

	patientDTO := &patient.PatientDTO{
		Name: "Jane Doe",
		Age:  29,
	}
	body, _ := json.Marshal(patientDTO)

	req, _ := http.NewRequest(http.MethodPut, "/patients/"+id.String(), bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Jane Doe")
}

func TestDeletePatient(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	patientService := &mockPatientService{}
	handler := handler.NewPatientsHandler(patientService)

	id := uuid.New()
	r.DELETE("/patients/:id", handler.DeletePatient)

	req, _ := http.NewRequest(http.MethodDelete, "/patients/"+id.String(), nil)
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}
