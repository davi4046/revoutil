package revoutil

import (
	"testing"
)

func TestNewLightKey(t *testing.T) {
	key := NewLightKey(10, 2741)
	t.Error(key)
}
