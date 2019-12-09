package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuccessfulLookup(t *testing.T) {
	assert := assert.New(t)

	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	var args = "foo.go --network 81.2.69.142 --db ../../test/data/test-data/GeoIP2-City-Test.mmdb"
	os.Args = strings.Split(args, " ")

	main()

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	assert.Contains(fmt.Sprintf("%s", out), "London")
}
