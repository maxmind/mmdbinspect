// mmdbinspect looks up records for one or more IPs/networks in one or more
// .mmdb databases
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type arrayFlags []string

func (i *arrayFlags) String() string {
	return strings.Join(*i, ", ")
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

func main() {
	var mmdb arrayFlags

	flag.Var(&mmdb, "db", "Path to an mmdb file. You may pass this arg more than once.")

	includeAliasedNetworks := flag.Bool(
		"include-aliased-networks",
		false,
		"Include aliased networks (e.g. 6to4, Teredo). This option may cause IPv4 networks to be listed more than once via aliases.",
	)

	includeBuildTime := flag.Bool(
		"include-build-time",
		false,
		"Include the build time of the database in the output.",
	)

	includeNetworksWithoutData := flag.Bool(
		"include-networks-without-data",
		false,
		`Include networks that have no data in the database. The "record" will be null for these.`,
	)

	includeEmptyValues := flag.Bool(
		"include-empty-values",
		false,
		"Include networks whose data is an empty map or empty array.",
	)

	useJSONL := flag.Bool("jsonl", false, "Output as JSONL instead of YAML.")

	flag.Usage = usage
	flag.Parse()

	// Any remaining arguments (not passed via flags) should be networks
	networks := flag.Args()

	if len(networks) == 0 {
		fmt.Println("You must provide at least one network address")
		usage()
		os.Exit(1)
	}

	if len(mmdb) == 0 {
		fmt.Println("You must provide a path to at least one .mmdb file")
		usage()
		os.Exit(1)
	}

	w := os.Stdout

	err := process(
		w,
		*useJSONL,
		networks,
		mmdb,
		*includeAliasedNetworks,
		*includeBuildTime,
		*includeNetworksWithoutData,
		*includeEmptyValues,
	)
	if err != nil {
		log.Fatal(err)
	}
}

func usage() {
	fmt.Printf(
		"Usage: %s [-include-aliased-networks] -db path/to/db -db path/to/other/db 130.113.64.30/24 0:0:0:0:0:ffff:8064:a678\n",
		os.Args[0],
	)
	flag.PrintDefaults()
	fmt.Print(`
The -db parameter may be a path to an MMDB file or a glob matching one or more
MMDB files.

Any additional arguments passed are assumed to be networks to look up. If an
address range is not supplied, /32 will be assumed for ipv4 addresses and /128
will be assumed for ipv6 addresses.
`)
}
