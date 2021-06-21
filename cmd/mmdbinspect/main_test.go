package main

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSuccessfulLookup(t *testing.T) {
	assert := assert.New(t)

	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	args := "foo.go --db ../../test/data/test-data/GeoIP2-City-Test.mmdb 81.2.69.142"
	os.Args = strings.Split(args, " ")

	main()

	require.NoError(t, w.Close())
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	assert.Contains(string(out), "London")
}
