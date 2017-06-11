package internal_test

import (
	"github.com/maprost/playground/pingserver/internal/test"
	"testing"
)

func TestPing(t *testing.T) {
	assert := test.InitRestTest(t)

	test.Ping204(assert)
}
