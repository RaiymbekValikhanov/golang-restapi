package model_test

import (
	"testing"

	"github.com/RaiymbekValikhanov/golang-restapi/internal/app/model"
	"github.com/stretchr/testify/assert"
)

func TestUser_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		u       func() *model.User
		isValid bool
	}{
		{
			name: "valid",
			u: func() *model.User {
				return model.TestUser(t)
			},
			isValid: true,
		},
		{
			name: "empty email",
			u: func() *model.User {
				return &model.User{
					Email: "",
					Password: "password",
				}
			},
			isValid: false,
		},
		{
			name: "invalid email",
			u: func() *model.User {
				return &model.User{
					Email: "invalid",
					Password: "password",
				}
			},
			isValid: false,
		},
		{
			name: "invalid password",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Password = "123"
				return u
			},
			isValid: false,
		},
		{
			name: "with enc password",
			u: func() *model.User {
				u := model.TestUser(t)
				u.Password = ""
				u.EncryptedPassword = "encryptedpassword"
				return u
			},
			isValid: true,
		},

	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.u().Validate())
			} else {
				assert.Error(t, tc.u().Validate())
			}
		})
	}
}

func TestUser_BeforeCreate(t *testing.T) {
	u := model.TestUser(t)

	assert.NoError(t, u.BeforeCreate())
	assert.NotEmpty(t, u.EncryptedPassword)
}
