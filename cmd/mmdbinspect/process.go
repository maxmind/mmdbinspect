package main

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/goccy/go-yaml"
)

func process(
	w io.Writer,
	useJSONL bool,
	networks, databases []string,
	includeAliasedNetworks,
	includeBuildTime,
	includeNetworksWithoutData bool,
) error {
	var encoder interface {
		Encode(any) error
	}

	if useJSONL {
		enc := json.NewEncoder(w)
		enc.SetEscapeHTML(false) // don't escape ampersands and angle brackets
		encoder = enc
	} else {
		encoder = yaml.NewEncoder(w)
	}

	iterator := records(
		networks,
		databases,
		includeAliasedNetworks,
		includeBuildTime,
		includeNetworksWithoutData,
	)

	for r, err := range iterator {
		if err != nil {
			return err
		}

		if err := encoder.Encode(r); err != nil {
			return fmt.Errorf("encoding record: %w", err)
		}
	}

	return nil
}
