package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/spf13/cobra"

	"github.com/desmos-labs/changeset/types"
	"github.com/desmos-labs/changeset/utils"
)

const (
	flagPath    = "path"
	flagDryRun  = "dry-run"
	flagVersion = "version"
)

// CollectCmd returns the Cobra command allowing to collect the changeset entries into a CHANGELOG entry
func CollectCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "collect",
		Short: "Collect the current changeset entries into a new version.",
		Long: `Collect the current changeset entries and create a Changelog entry.
If a version is specified, that version will be used as the next version. Otherwise, the next version
will be computed considering the changeset entries and their types following the semantic versioning specification.`,
		Args:    cobra.NoArgs,
		PreRunE: readConfig,
		RunE:    runCollect,
	}

	curDir, _ := os.Getwd()
	cmd.Flags().String(flagPath, path.Join(curDir, "CHANGELOG.md"), "Path to the CHANGELOG file")
	cmd.Flags().Bool(flagDryRun, false, "Perform a simulation of the changes, without updating the file")
	cmd.Flags().String(flagVersion, "", "Version to be used")

	return cmd
}

func runCollect(cmd *cobra.Command, args []string) error {
	entries, err := utils.GetEntries()
	if err != nil {
		return err
	}

	// Get the version
	var version *types.Version
	ver, _ := cmd.Flags().GetString(flagVersion)
	if ver != "" {
		// The user has specified a custom version, so we use that
		v, err := types.ParseVersion(ver)
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

	// Crete the changelog entry
	changes := utils.CollectChanges(entries)
	changelog := types.NewChangeLog(version, changes)

	// Convert the entry to markdown
	out, err := utils.ConvertToMarkdown(cfg, changelog)
	if err != nil {
		return err
	}

	// Update the current changelog
	changelogPath, _ := cmd.Flags().GetString(flagPath)
	updated, err := utils.UpdateChangelog(out, changelogPath)
	if err != nil {
		return err
	}

	dryRun, _ := cmd.Flags().GetBool(flagDryRun)
	if dryRun {
		// If dry-run simply output the changes
		cmd.Println(updated)
		return nil
	}

	// Write the file
	err = ioutil.WriteFile(changelogPath, []byte(updated), 0666)
	if err != nil {
		return fmt.Errorf("error while writing changelog file: %s", err)
	}

	return nil
}
