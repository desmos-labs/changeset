package cmd

import "github.com/spf13/cobra"

// RootCmd returns the root command of the application
func RootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "changeset",
		Short: "Changeset is a tool that allows to easily create new changelog entries and releave new versions",
	}

	cmd.AddCommand(
		InitCmd(),
		AddCmd(),
		CollectCmd(),
	)

	return cmd
}
