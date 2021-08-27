package store

import (
	"fmt"
	"strings"
	"testing"
)

func TestStore(t *testing.T, databaseurl string) (*Store, func (...string)) {
	t.Helper()

	config := NewConfig()
	config.DataBaseURL = databaseurl
	store := NewStore(config)

	if err := store.Open(); err != nil {
		t.Fatal(err) 
	}

	return store, func(tables ...string) {
		if len(tables) > 0 {
			if _, err := store.db.Exec(
				fmt.Sprintf("TRUNCATE %s CASCADE", strings.Join(tables, ", ")),
			); err != nil {
				t.Fatal(err)
			}
		}
		store.Close()
	}
}