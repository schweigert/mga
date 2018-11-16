package randomizer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInt(t *testing.T) {
	value := Int(3)
	assert.True(t, value == 0 || value == 1 || value == 2)
	assert.Zero(t, Int(0))
}

func BenchmarkInt(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Int(n)
	}
}
