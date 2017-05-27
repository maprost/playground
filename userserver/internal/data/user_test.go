package data_test

import (
	"github.com/maprost/playground/userserver/internal/data"
	"github.com/maprost/playground/userserver/internal/test"
	"testing"
)

func TestUserWorkflow(t *testing.T) {
	assert := test.InitUnitTest(t)

	user := data.UserAccount{
		Email:        "user@DataTestUserWorkflow.de",
		PasswordHash: "12345678901234567890",
	}
	err := user.Insert()
	assert.Nil(err)

}
