package cmd

import (
	"github.com/spf13/cobra"

	"github.com/desmos-labs/changeset/types"
	"github.com/desmos-labs/changeset/utils"
)

// InitCmd returns the Cobra command that allows to initialize the current folder to use changeset
func InitCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "init [github-repo] [version]",
		Args:  cobra.ExactArgs(2),
		Short: "Initialize the current folder to use changeset starting from the current version",
		RunE: func(cmd *cobra.Command, args []string) error {
			version, err := types.ParseVersion(args[1])
			if err != nil {
				return err
			}

			config := types.DefaultConfig(args[0], version)
			err = config.Validate()
			if err != nil {
				return err
			}

			return utils.WriteConfig(config)
		},
	}
}
