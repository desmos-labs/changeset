package utils

import (
	"fmt"
	"github.com/desmos-labs/changeset/types"
	"strings"
)

// CollectChanges collects the given entries as a types.TypeChanges instance
func CollectChanges(entries []*types.Entry) types.TypeChanges {
	var typesChanges = types.TypeChanges{}
	for _, entry := range entries {
		existingChanges := typesChanges[entry.Type]
		if existingChanges == nil {
			existingChanges = types.ModuleChanges{}
		}

		existingChanges[entry.Module] = append(existingChanges[entry.Module], entry)
		typesChanges[entry.Type] = existingChanges
	}

	return typesChanges
}

// ConvertToMarkdown converts the given changelog to a Markdown representation
func ConvertToMarkdown(config *types.Config, changelog *types.ChangeLog) (string, error) {
	output := fmt.Sprintf("## %s\n", changelog.Version)

	for typeCode, changes := range changelog.Changes {
		changesType, err := config.GetTypeByCode(typeCode)
		if err != nil {
			return "", err
		}

		output += fmt.Sprintf("### %s\n", strings.Title(changesType.Code.String()))
		for moduleID, entries := range changes {
			module, err := config.GetModuleByID(moduleID)
			if err != nil {
				return "", err
			}

			output += fmt.Sprintf("#### %s\n", module.Description)
			for _, entry := range entries {
				output += fmt.Sprintf("* ([%[1]d](%[2]s/pull/%[1]d)) %[3]s\n",
					entry.PullRequestID, config.GitHubRepo, entry.Description)
			}
			output += "\n"
		}
	}

	return output, nil
}
