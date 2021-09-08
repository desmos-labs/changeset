package cmd

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"

	"github.com/desmos-labs/changeset/types"
	"github.com/desmos-labs/changeset/utils"
)

// AddCmd returns the Cobra command that allows to add a new changeset entry
func AddCmd() *cobra.Command {
	return &cobra.Command{
		Use:     "add",
		Short:   "Create a new change entry for the next version",
		PreRunE: readConfig,
		RunE:    runAdd,
	}
}

// runAdd represents the function that is run when using the add command
func runAdd(cmd *cobra.Command, _ []string) error {
	selectedType, err := selectType()
	if err != nil {
		return err
	}

	isBackwardCompatible, err := selectBackWardCompatible()
	if err != nil {
		return err
	}

	selectedModule, err := selectModule()
	if err != nil {
		return err
	}

	getPRID, err := getPullRequestID()
	if err != nil {
		return err
	}

	description, err := getDescription()
	if err != nil {
		return err
	}

	err = utils.WriteEntry(types.NewEntry(
		selectedType.Code,
		selectedModule.Code,
		getPRID,
		description,
		isBackwardCompatible,
		time.Now().UTC(),
	))
	if err != nil {
		return err
	}

	cmd.Println("\nEntry added successfully. Make sure to add it to Git")
	return nil
}

// selectType allows to select the type of the change
func selectType() (*types.Type, error) {
	prompt := promptui.Select{
		Label: "Select the type of this change",
		Items: cfg.Types,

		Templates: &promptui.SelectTemplates{
			Active:   "\U00002713 {{ .Description | cyan }}",
			Inactive: "  {{ .Description | cyan }}",
			Selected: "Change type: \U00002713 {{ .Description | cyan }}",
		},
	}

	index, _, err := prompt.Run()
	if err != nil {
		return nil, err
	}

	return cfg.Types[index], nil
}

// selectBackWardCompatible allows to set whether the change is backward compatible or not
func selectBackWardCompatible() (bool, error) {
	prompt := promptui.Select{
		Label: "Is this change backward compatible",
		Items: []string{"Yes", "No"},

		Templates: &promptui.SelectTemplates{
			Active:   "\U00002713 {{ . | cyan }}",
			Inactive: "  {{ . | cyan }}",
			Selected: "Backward compatible: \U00002713 {{ . | cyan }}",
		},
	}

	index, _, err := prompt.Run()
	if err != nil {
		return false, err
	}

	return index == 0, nil
}

// selectModule allows to select which module this change refers to
func selectModule() (*types.Module, error) {
	if len(cfg.Modules) == 0 {
		return types.ModuleNone, nil
	}

	prompt := promptui.Select{
		Label: "Module",
		Items: cfg.Modules,

		Templates: &promptui.SelectTemplates{
			Active:   "\U00002713 {{ .Description | cyan }}",
			Inactive: "  {{ .Description | cyan }}",
			Selected: "Module: \U00002713 {{ .Description | cyan }}",
		},
	}

	index, _, err := prompt.Run()
	if err != nil {
		return nil, err
	}

	return cfg.Modules[index], nil
}

// getDescription allows the user to set a description for the change
func getDescription() (string, error) {
	prompt := promptui.Prompt{
		Label: "Please write a brief description of the change",
		Validate: func(s string) error {
			if strings.TrimSpace(s) == "" {
				return fmt.Errorf("Invalid description")
			}
			return nil
		},
	}
	return prompt.Run()
}

// getPullRequestID allows to get the number of the PR to which this change is associated
func getPullRequestID() (int, error) {
	prompt := promptui.Prompt{
		Label: "What is the number of the PR associated to this change",
		Validate: func(s string) error {
			if s == "" {
				return nil
			}

			value, err := strconv.Atoi(s)
			if err != nil || value <= 0 {
				return fmt.Errorf("Invalid PR number")
			}
			return nil
		},
	}

	value, err := prompt.Run()
	if err != nil {
		return -1, err
	}

	return strconv.Atoi(value)
}
