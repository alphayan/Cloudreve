package serializer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRequestSignString(t *testing.T) {
	asserts := assert.New(t)

	sign := NewRequestSignString("1", "2", "3")
	asserts.NotEmpty(sign)
}
