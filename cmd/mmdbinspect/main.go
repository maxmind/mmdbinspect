package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/oalders/mmdbinspect/pkg/mmdbinspect"
)

type arrayFlags []string

func (i *arrayFlags) String() string {
	return "my string representation"
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

func main() {
	var mmdb, network arrayFlags

	flag.Var(&mmdb, "db", "Path to an mmdb file. You may pass this arg more than once.")
	flag.Var(&network, "network", "A network address to look up. You may pass this arg more than once.")

	flag.Parse()

	if len(network) == 0 {
		fmt.Println("You must provide at least one network address")
		flag.PrintDefaults()
		os.Exit(1)
	}

	if len(mmdb) == 0 {
		fmt.Println("You must provide a path to at least one .mmdb file")
		flag.PrintDefaults()
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

	fmt.Printf("%v", json)
}
