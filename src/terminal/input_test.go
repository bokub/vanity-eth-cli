package terminal

import (
	"bufio"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadBool(t *testing.T) {
	val, err := ReadBool(bufio.NewReader(strings.NewReader("n\n")))
	assert.Nil(t, err)
	assert.False(t, val)

	val, err = ReadBool(bufio.NewReader(strings.NewReader("y\n")))
	assert.Nil(t, err)
	assert.True(t, val)
}

func TestReadString(t *testing.T) {
	val, err := ReadString(bufio.NewReader(strings.NewReader("Abc\n")))
	assert.Nil(t, err)
	assert.Equal(t, "Abc", val)

	val, err = ReadString(bufio.NewReader(strings.NewReader("Abx\n")))
	assert.NotNil(t, err)
}
