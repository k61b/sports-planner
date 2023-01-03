package exercises

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/kayraberktuncer/sports-planner/pkg/common/models"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockDB struct {
	mock.Mock
	models.Store
}

func (m *MockDB) Find(value interface{}, where ...interface{}) *gorm.DB {
	args := m.Called(value)
	return args.Get(0).(*gorm.DB)
}

func (m *MockDB) Error() error {
	args := m.Called()
	return args.Error(0)
}

func TestGetExercises(t *testing.T) {
	// Setup mock DB
	mockDB := new(MockDB)
	mockDB.On("Find", &[]models.Exercise{}).Return(mockDB).Once()

	// Setup Gin router
	router := gin.New()
	router.GET("/", func(c *gin.Context) {
		handlers := &Handlers{DB: mockDB}
		handlers.GetExercises(c)
	})

	// Perform request
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	// Assert response
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	// Assert mock DB was called correctly
	mockDB.AssertExpectations(t)
}
