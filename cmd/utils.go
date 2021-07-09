package cmd

import (
	"github.com/desmos-labs/changeset/types"
	"github.com/desmos-labs/changeset/utils"
	"github.com/spf13/cobra"
)

// cfg represents the configuration that has been parsed
var cfg *types.Config

// readConfig allows to read the current configuration
func readConfig(_ *cobra.Command, _ []string) error {
	config, err := utils.ReadConfig()
	if err != nil {
		return err
	}

	cfg = config
	return nil
}
