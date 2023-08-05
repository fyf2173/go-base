package common

import (
	"testing"

	"github.com/google/uuid"
)

func TestNewUUid(t *testing.T) {
	for i := 0; i <= 2; i++ {
		t.Log(uuid.New())
	}
}
