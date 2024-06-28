package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_login_with_admin(t *testing.T) {
	_, err := Login("admin@admin.com", "password123")

	assert.NoError(t, err)
}

func Test_failed_login_with_admin(t *testing.T) {
	_, err := Login("admin@admin.com", "password12")

	assert.Error(t, err)
}
