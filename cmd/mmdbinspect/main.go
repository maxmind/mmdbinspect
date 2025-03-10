// mmdbinspect looks up records for one or more IPs/networks in one or more
// .mmdb databases
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/maxmind/mmdbinspect/v2/pkg/mmdbinspect"
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
		"Include aliased networks (e.g. 6to4, Teredo). This option may cause IPv4 networks to be listed more than once via aliases.", //nolint: lll
	)

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

	encoder := json.NewEncoder(os.Stdout)

	iterator := mmdbinspect.Records(network, mmdb, *includeAliasedNetworks)

	for r, err := range iterator {
		if err != nil {
			log.Fatal(err)
		}

		err = encoder.Encode(r)
		if err != nil {
			log.Fatal(fmt.Errorf("encoding record: %w", err))
		}
	}
}

func usage() {
	fmt.Printf(
		"Usage: %s [-include-aliased-networks] -db path/to/db -db path/to/other/db 130.113.64.30/24 0:0:0:0:0:ffff:8064:a678\n", //nolint: lll
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
