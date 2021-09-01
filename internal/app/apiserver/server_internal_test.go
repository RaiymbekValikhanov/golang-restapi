package apiserver

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/RaiymbekValikhanov/golang-restapi/internal/app/model"
	"github.com/RaiymbekValikhanov/golang-restapi/internal/app/store/teststore"
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
