package model

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"log"
)

func TestUser(t *testing.T) {
	var id int = 1
	var name string = "hoge"
	var email string = "hoge@hoge.hoge"
	var password string = "password"
	// user := NewUser(id, name, email, password)

	var miss_id int = 2
	var miss_name string = "fuga"
	var miss_email string = "fuga@fuga.fuga"
	var miss_password string = "PassWord"
	miss_user := NewUser(miss_id, miss_name, miss_email, miss_password)

	// if user == nil {
	// 	t.Errorf("failed NewUser()")
	// }

	if miss_user == nil {
		t.Errorf("failed NewUser()")
	}

	log.Print("Here, pass test")

	// assert.Equal(t, id, user.ID)
	// assert.Equal(t, name, user.Name)
	// assert.Equal(t, email, user.Email)
	// assert.Equal(t, password, user.Password)

	// t.Logf("user: %p", user)
	// t.Logf("user.ID: %d", user.ID)
	// t.Logf("user.Name: %s", user.Name)
	// t.Logf("user.Email: %s", user.Email)
	// t.Logf("user.Password: %s", user.Password)

	log.Print("Here, miss test")

	assert.Equal(t, id, miss_user.ID)
	assert.Equal(t, name, miss_user.Name)
	assert.Equal(t, email, miss_user.Email)
	assert.Equal(t, password, miss_user.Password)
}
