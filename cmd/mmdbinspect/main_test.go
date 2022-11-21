package main

import (
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSuccessfulLookup(t *testing.T) {
	a := assert.New(t)

	rescueStdout := os.Stdout
	r, w, err := os.Pipe()
	require.NoError(t, err)
	os.Stdout = w

	args := "foo.go -db ../../test/data/test-data/GeoIP2-City-Test.mmdb 81.2.69.142"
	os.Args = strings.Split(args, " ")

	main()

	require.NoError(t, w.Close())
	out, err := io.ReadAll(r)
	require.NoError(t, err)
	os.Stdout = rescueStdout

	a.Contains(string(out), "London")
}
