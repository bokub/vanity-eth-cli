package main

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainProgram(t *testing.T) {
	t.Skip() // TODO perform test
	return

	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	outC := make(chan string)
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	w.Close()
	os.Stdout = oldStdout
	out := <-outC

	assert.Contains(t, out, "Address found")
	assert.Contains(t, out, "Address:     0xAb0")
}
