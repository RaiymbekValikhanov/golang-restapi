package apiserver

import (
	"bytes"
	"encoding/json"
	"fmt"

	// fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/RaiymbekValikhanov/golang-restapi/internal/app/model"
	"github.com/RaiymbekValikhanov/golang-restapi/internal/app/store/teststore"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/stretchr/testify/assert"
)

func TestServer_HandleUsersCreate(t *testing.T) {
	s := NewServer(teststore.NewStore(), sessions.NewCookieStore([]byte("secret")))

	testCases := []struct {
		name       string
		payload    interface{}
		statusCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"email":    "user@example.org",
				"password": "password",
			},
			statusCode: http.StatusCreated,
		},
		{
			name:       "invalid payload",
			payload:    "invalid",
			statusCode: http.StatusBadRequest,
		},
		{
			name: "invalid params",
			payload: map[string]string{
				"email": "invalid",
			},
			statusCode: http.StatusUnprocessableEntity,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/users", b)

			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.statusCode, rec.Code)
		})
	}
}

func TestServer_HandleSessionCreate(t *testing.T) {
	store := teststore.NewStore()
	u := model.TestUser(t)
	store.User().Create(u)
	s := NewServer(store, sessions.NewCookieStore([]byte("secret")))

	testCases := []struct {
		name       string
		payload    interface{}
		statusCode int
	}{
		{
			name: "valid",
			payload: map[string]string{
				"email":    "user@example.org",
				"password": "password",
			},
			statusCode: http.StatusOK,
		},
		{
			name: "invalid payload",
			payload: map[string]string{
				"email":    "kek",
				"password": "",
			},
			statusCode: http.StatusUnauthorized,
		},
		{
			name: "invalid password",
			payload: map[string]string{
				"email":    "user@example.org",
				"password": "123132123",
			},
			statusCode: http.StatusUnauthorized,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			b := &bytes.Buffer{}
			json.NewEncoder(b).Encode(tc.payload)
			req, _ := http.NewRequest(http.MethodPost, "/sessions", b)

			s.ServeHTTP(rec, req)
			assert.Equal(t, tc.statusCode, rec.Code)
		})
	}
	
}

func TestServer_AuthenticateUser(t *testing.T) {
	store := teststore.NewStore()
	u := model.TestUser(t)
	store.User().Create(u)

	testCases := []struct {
		name        string
		cookieValue map[interface{}]interface{}
		statusCode  int
	}{
		{
			name: "valid",
			cookieValue: map[interface{}]interface{}{
				"user_id": u.ID,
			},
			statusCode: http.StatusOK,
		}, 
		{
			name: "no auth",
			cookieValue: nil,
			statusCode: http.StatusUnauthorized,
		},
	}

	
	s := NewServer(store, sessions.NewCookieStore([]byte("secret")))
	sc := securecookie.New([]byte("secret"), nil)
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rec := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodPost, "/", nil)
			cookieStr, _ := sc.Encode(sessionName, tc.cookieValue)
			req.Header.Set("Cookie", fmt.Sprintf("%s=%s", sessionName, cookieStr))
			s.authenticateUser(handler).ServeHTTP(rec, req)
			assert.Equal(t, tc.statusCode, rec.Code)
		})
	}
}

