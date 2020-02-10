package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/maxmind/mmdbinspect/pkg/mmdbinspect"
)

type arrayFlags []string

func (i *arrayFlags) String() string {
	return strings.Join(*i, ", ")
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

func usage() {
	fmt.Printf(
		"Usage: %s --db path/to/db --db path/to/other/db 130.113.64.30/24 0:0:0:0:0:ffff:8064:a678\n",
		os.Args[0],
	)
	flag.PrintDefaults()
	fmt.Print(`
Any additional arguments passed are assumed to be networks to look up.  If an
address range is not supplied, /32 will be assumed for ipv4 addresses and /128
will be assumed for ipv6 addresses.
`)
}

func main() {
	var mmdb arrayFlags

	flag.Var(&mmdb, "db", "Path to an mmdb file. You may pass this arg more than once.")

	flag.Usage = usage
	flag.Parse()

	// Any remaining arguments (not passed via flags) should be networks
	network := flag.Args()

	if len(network) == 0 {
		fmt.Println("You must provide at least one network address")
		usage()
		os.Exit(1)
	}

	if len(mmdb) == 0 {
		fmt.Println("You must provide a path to at least one .mmdb file")
		usage()
		os.Exit(1)
	}

	records, err := mmdbinspect.AggregatedRecords(network, mmdb)

	if err != nil {
		log.Fatal(err)
	}

	json, err := mmdbinspect.RecordToString(records)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", json)
}
