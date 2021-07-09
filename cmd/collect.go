package cmd

import (
	"github.com/spf13/cobra"

	"github.com/desmos-labs/changeset/types"
	"github.com/desmos-labs/changeset/utils"
)

// CollectCmd returns the Cobra command allowing to collect the changeset entries into a CHANGELOG entry
func CollectCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "collect [[version]]",
		Short: "Collect the current changeset entries into a new version.",
		Long: `
Collect the current changeset entries and create a Changelog entry.
If a version is specified, that version will be used as the next version. Otherwise, the next version
will be computed considering the changeset entries and their types following the semantic versioning specification.`,
		Args:    cobra.RangeArgs(0, 1),
		PreRunE: readConfig,
		RunE:    runCollect,
	}
}

func runCollect(cmd *cobra.Command, args []string) error {
	entries, err := utils.GetEntries()
	if err != nil {
		return err
	}

	// Get the version
	var version *types.Version
	if len(args) == 1 {
		// The user has specified a custom version, so we use that
		v, err := types.ParseVersion(args[0])
		if err != nil {
			return err
		}
		version = v
	} else {
		// The user has not specified a custom version, so we need to compute the next one
		v, err := types.GetNextVersion(cfg, entries)
		if err != nil {
			return err
		}
		version = v
	}

	changes := utils.CollectChanges(entries)
	changelog := types.NewChangeLog(version, changes)

	out, err := utils.ConvertToMarkdown(cfg, changelog)
	if err != nil {
		return err
	}

	cmd.Println(out)
	return nil
}
