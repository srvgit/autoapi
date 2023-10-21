package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateUUID(t *testing.T) {
	uuid1 := GenerateUUID()
	assert.NotEmpty(t, uuid1, "UUID should not be empty")

	uuid2 := GenerateUUID()
	assert.NotEqual(t, uuid1, uuid2, "Generated UUIDs should be unique")
}
