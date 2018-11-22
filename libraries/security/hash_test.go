package security

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHash(t *testing.T) {
	value := "zelda"
	assert.Equal(t, "5a2fa4da9967553d347c13a61017f93facfcc025", Hash(value))
}

func BenchmarkHash(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Hash("zelda")
	}
}
