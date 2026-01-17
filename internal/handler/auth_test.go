package handler

import (

	"errors"
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

func TestLogin_Failure(t *testing.T){
	gin.SetMode(gin.TestMode)

	mockAuth := &mocks.AuthServiceMock{
		LoginFunc: func(username,password string)(string,error){
			return "", errors.New("invalid credentials")
		},
	}

	h := NewHandler(nil,mockAuth)

	router := gin.New()

	router.POST("/auth/login", h.login)

	body := `{"username":"admin","password":"213das"}`

	req := httptest.NewRequest(http.MethodPost,"/auth/login",strings.NewReader(body))

	req.Header.Set("Content-Type","application/json")
	w := httptest.NewRecorder()


	router.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized{
		t.Fatalf("expected 401, got : %d",w.Code)
	}

}

func TestRegister_Success(t *testing.T){
	gin.SetMode(gin.TestMode)

	mockAuth := &mocks.AuthServiceMock{
		RegisterFunc: func (username,password string)error  {
			return nil
		},
	}

	h := NewHandler(nil,mockAuth)

	router := gin.New()

	router.POST("/auth/register",h.register)

	body := `{"username":"newuser","password":"newpassword"}`

	req := httptest.NewRequest(http.MethodPost,"/auth/register",strings.NewReader(body))
	req.Header.Set("Content-Type","application/json")

	w := httptest.NewRecorder()

	router.ServeHTTP(w,req)

	if w.Code != http.StatusOK{
		t.Fatalf("expected 200, got : %d",w.Code)
	}

}

func TestRegister_Failure(t *testing.T){
	gin.SetMode(gin.TestMode)


	mockAuth := &mocks.AuthServiceMock{
		RegisterFunc: func (username, password string)error {
			return errors.New("User Already exist")
		},
	}
	h := NewHandler(nil, mockAuth)
	router := gin.New()
	router.POST("/auth/register", h.register)

	body := `{"username":"existinuser","password":"password123"}`
	req := httptest.NewRequest(http.MethodPost,"/auth/register",strings.NewReader(body))

	req.Header.Set("Content-Type","application/json")

	w := httptest.NewRecorder()


	router.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest{
		t.Fatalf("expected 400, got : %d",w.Code)
	}

}