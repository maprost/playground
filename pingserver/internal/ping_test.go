package internal_test

import (
	"testing"

	"github.com/maprost/playground/pingserver/internal/test"
)

func TestPing(t *testing.T) {
	assert := test.InitRestTest(t)

	test.Ping200(assert)
}
