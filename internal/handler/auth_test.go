package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"


	"github.com/Coldwws/todo/internal/handler/mocks"
	"github.com/gin-gonic/gin"
)


func TestLogin_Success(t *testing.T){
	gin.SetMode(gin.TestMode)

	// Создаем мок для нашего сервиса

	mockAuth := &mocks.AuthServiceMock{
		LoginFunc: func(name,password string) (string,error){
			return "mocked_token",nil
		},
	}

	// Здесь будет handler  для мока

	h := NewHandler(nil, mockAuth) // в основном Handler struct есть 2 сервиса, нам нужен только auth поэтому nil for first
	
	router := gin.New()
	router.POST("/auth/login", h.login)

	body := `{"username":"testuser", "password":"testpassword"}`
	req := httptest.NewRequest(http.MethodPost, "/auth/login", strings.NewReader(body))
	req.Header.Set("Content-Type","application/json")

	w := httptest.NewRecorder()


	router.ServeHTTP(w, req)


	if w.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", w.Code)
	}

	if !strings.Contains(w.Body.String(), "token") {
		t.Fatalf("expected token in response, got %s", w.Body.String())
	}
}