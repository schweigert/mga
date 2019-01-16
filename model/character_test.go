package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMoveCharacter(t *testing.T) {
	character := Character{PositionX: 5, PositionY: 5, DirectionX: -1, DirectionY: -1}
	character.Move()
	assert.Equal(t, 4, character.PositionX)
	assert.Equal(t, 4, character.PositionY)
	assert.Equal(t, 0, character.DirectionX, character.DirectionY)
}
