package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	handler "github.com/slilp/go-auth/handler/account"
	mocks "github.com/slilp/go-auth/mocks/mock_service"
	service "github.com/slilp/go-auth/service/account"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() (*gin.Context, *gin.Engine) {
	gin.SetMode(gin.TestMode)
	// router := gin.Default()
	w := httptest.NewRecorder()
	ctx, router := gin.CreateTestContext(w)

	return ctx, router
}

func newFakeAccount() service.AccountInfo {
	return service.AccountInfo{
		Username:  "faleUsername",
		FirstName: "fakeFirstName",
		LastName:  "fakeLastName",
		Avatar:    "fakeAvatar",
		Role:      "fakeRole",
	}
}

func newFakeCreateAccount() service.CreateAccountDto {
	return service.CreateAccountDto{
		Username:  "faleUsername",
		Password:  "fakePassword",
		FirstName: "fakeFirstName",
		LastName:  "fakeLastName",
		Avatar:    "fakeAvatar",
	}
}

func TestRegister(t *testing.T) {

	t.Run("Error invalid request body", func(t *testing.T) {

		//Arrange
		ctrl := gomock.NewController(t)
		mockAccountService := mocks.NewMockAccountService(ctrl)

		c, r := SetUpRouter()
		accHandler := handler.NewAccountHttpHandler(mockAccountService)
		r.POST("/auth/register", accHandler.Register)
		reqBody := service.CreateAccountDto{}
		jsonReq, _ := json.Marshal(reqBody)

		//Act
		c.Request, _ = http.NewRequest("POST", "/auth/register", bytes.NewBuffer(jsonReq))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, c.Request)

		//Assert
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Internal server error create new account", func(t *testing.T) {

		//Arrange
		ctrl := gomock.NewController(t)
		mockAccountService := mocks.NewMockAccountService(ctrl)
		mockAccountService.EXPECT().CreateAccount(gomock.Any()).Return(nil, assert.AnError)

		c, r := SetUpRouter()
		accHandler := handler.NewAccountHttpHandler(mockAccountService)
		r.POST("/auth/register", accHandler.Register)
		reqBody := newFakeCreateAccount()
		jsonReq, _ := json.Marshal(reqBody)

		//Act
		c.Request, _ = http.NewRequest("POST", "/auth/register", bytes.NewBuffer(jsonReq))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, c.Request)

		//Assert
		assert.Equal(t, http.StatusInternalServerError, w.Code)
	})

	t.Run("Success", func(t *testing.T) {
		//Arrange
		ctrl := gomock.NewController(t)
		mockAccountService := mocks.NewMockAccountService(ctrl)
		fakeAccount := newFakeAccount()

		mockAccountService.EXPECT().CreateAccount(gomock.Any()).Return(&fakeAccount, nil)

		c, r := SetUpRouter()
		accHandler := handler.NewAccountHttpHandler(mockAccountService)
		r.POST("/auth/register", accHandler.Register)
		reqBody := newFakeCreateAccount()
		jsonReq, _ := json.Marshal(reqBody)

		//Act
		c.Request, _ = http.NewRequest("POST", "/auth/register", bytes.NewBuffer(jsonReq))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, c.Request)
		var accRes service.AccountInfo
		json.Unmarshal(w.Body.Bytes(), &accRes)

		//Assert
		assert.Equal(t, http.StatusCreated, w.Code)
		assert.Equal(t, fakeAccount, accRes)

	})

}
