package main

import (
	"flag"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSuccessfulLookup(t *testing.T) {
	a := assert.New(t)

	// Reset flag package for this test
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

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

func TestVersionFlag(t *testing.T) {
	a := assert.New(t)

	// Reset flag package for this test
	flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)

	rescueStdout := os.Stdout
	r, w, err := os.Pipe()
	require.NoError(t, err)
	os.Stdout = w

	args := "mmdbinspect -version"
	os.Args = strings.Split(args, " ")

	func() {
		defer func() {
			if panicVal := recover(); panicVal != nil { //nolint:revive,staticcheck // Expected exit, ignore panic from os.Exit(0)
			}
		}()
		main()
	}()

	require.NoError(t, w.Close())
	out, err := io.ReadAll(r)
	require.NoError(t, err)
	os.Stdout = rescueStdout

	a.Equal("0.2.0\n", string(out))
}
