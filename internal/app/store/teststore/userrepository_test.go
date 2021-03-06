package teststore_test

import (
	"testing"

	"github.com/RaiymbekValikhanov/golang-restapi/internal/app/model"
	"github.com/RaiymbekValikhanov/golang-restapi/internal/app/store"
	"github.com/RaiymbekValikhanov/golang-restapi/internal/app/store/teststore"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	s := teststore.NewStore()
	u := model.TestUser(t)

	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	// Case 1: non-existent user
	// Expected: error

	s := teststore.NewStore()
	email := "user@example.org"
	_, err := s.User().FindByEmail(email)

	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	// Case 2: existent user
	// Expected: user
	s.User().Create(model.TestUser(t))
	u, err := s.User().FindByEmail(email)

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_Find(t *testing.T) {
	s := teststore.NewStore()
	u1 := model.TestUser(t)
	s.User().Create(u1)
	u2, err := s.User().Find(u1.ID)

	assert.NoError(t, err)
	assert.NotNil(t, u2)
}