package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestHasLetters(t *testing.T) {
	assert.True(t, HasLetters("123a456"))
	assert.True(t, HasLetters("0019A"))
	assert.False(t, HasLetters("01234765"))
}

func TestRoundDuration(t *testing.T) {
	assert.Equal(t, 1000*time.Millisecond, RoundDuration(1234*time.Millisecond, time.Second))
	assert.Equal(t, 1200*time.Millisecond, RoundDuration(1234*time.Millisecond, time.Second/10))
	assert.Equal(t, time.Minute, RoundDuration(59975*time.Millisecond, time.Second/10))
}
